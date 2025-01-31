package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Client struct
type S3Client struct {
	Client     *s3.Client
	Uploader   *manager.Uploader
	BucketName string
}

// NewS3Client initializes an S3 client using credentials from config
func NewS3Client() (*S3Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(s3Client)

	return &S3Client{
		Client:     s3Client,
		Uploader:   uploader,
		BucketName: os.Getenv("AWS_S3_BUCKET"),
	}, nil
}

// UploadFile uploads a file to S3
func (s *S3Client) UploadFile(file multipart.File, fileName string) (string, error) {
	_, err := s.Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		log.Printf("failed to upload file: %v", err)
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.BucketName, os.Getenv("AWS_REGION"), fileName), nil
}
