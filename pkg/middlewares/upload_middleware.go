package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

type contextKey string

const imageURLsKey contextKey = "image_urls"

type StorageType string

const (
	StorageS3    StorageType = "s3"
	StorageLocal StorageType = "local"
)

type ImageUploadMiddleware struct {
	S3Client    *s3.Client
	BucketName  string
	StoragePath string
	StorageType StorageType
	ServerURL   string
}

func NewImageUploadMiddleware(awsRegion, bucketName, storagePath string, storageType StorageType, serverURL string) (*ImageUploadMiddleware, error) {
	var s3Client *s3.Client

	if storageType == StorageS3 {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
		if err != nil {
			return nil, err
		}
		s3Client = s3.NewFromConfig(cfg)
	}

	if storageType == StorageLocal {
		if _, err := os.Stat(storagePath); os.IsNotExist(err) {
			err := os.MkdirAll(storagePath, os.ModePerm)
			if err != nil {
				return nil, fmt.Errorf("failed to create local storage directory: %v", err)
			}
		}
	}

	return &ImageUploadMiddleware{
		S3Client:    s3Client,
		BucketName:  bucketName,
		StoragePath: storagePath,
		StorageType: storageType,
		ServerURL:   serverURL,
	}, nil
}

func (m *ImageUploadMiddleware) UploadImage(serviceName, fieldName string, nextFunc func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		if r.MultipartForm == nil {
			http.Error(w, "Multipart form is missing", http.StatusBadRequest)
			return
		}

		files, ok := r.MultipartForm.File[fieldName]
		if !ok || len(files) == 0 {
			http.Error(w, "No file uploaded", http.StatusBadRequest)
			return
		}

		var uploadedURLs []string

		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, "Error reading file", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			if !contains([]string{".jpg", ".jpeg", ".png", ".gif"}, ext) {
				http.Error(w, "Invalid file type", http.StatusUnsupportedMediaType)
				return
			}

			fileName := fmt.Sprintf("%s/%s%s", serviceName, uuid.New().String(), ext)
			var fileURL string

			if m.StorageType == StorageS3 {
				uploader := manager.NewUploader(m.S3Client)
				_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
					Bucket: aws.String(m.BucketName),
					Key:    aws.String(fileName),
					Body:   file,
					ACL:    types.ObjectCannedACLPublicRead,
				})
				if err != nil {
					http.Error(w, "Error uploading file to S3", http.StatusInternalServerError)
					return
				}
				fileURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", m.BucketName, fileName)
			} else {
				localDir := filepath.Join(m.StoragePath, serviceName)
				if _, err := os.Stat(localDir); os.IsNotExist(err) {
					err := os.MkdirAll(localDir, os.ModePerm)
					if err != nil {
						http.Error(w, "Failed to create directory for local storage", http.StatusInternalServerError)
						return
					}
				}

				localFilePath := filepath.Join(m.StoragePath, fileName)
				outFile, err := os.Create(localFilePath)
				if err != nil {
					http.Error(w, "Error saving file locally", http.StatusInternalServerError)
					return
				}
				defer outFile.Close()

				if _, err := io.Copy(outFile, file); err != nil {
					http.Error(w, "Error writing file", http.StatusInternalServerError)
					return
				}

				fileURL = fmt.Sprintf("%s/uploads/%s", m.ServerURL, fileName)
			}

			uploadedURLs = append(uploadedURLs, fileURL)
		}

		ctx := context.WithValue(r.Context(), imageURLsKey, uploadedURLs)
		nextFunc(w, r.WithContext(ctx))
	})
}

func GetImageURLsFromContext(ctx context.Context) ([]string, bool) {
	urls, ok := ctx.Value(imageURLsKey).([]string)
	return urls, ok
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
