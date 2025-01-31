package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// UploadMiddleware handles file upload and saves the file to the server
func UploadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse the multipart form with a limit on file size (10 MB)
		err := r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			http.Error(w, "Unable to parse form: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Retrieve the file from the form field (field name is 'file')
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to retrieve file from form: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Generate a unique filename (could be dynamic based on user info, timestamp, etc.)
		fileName := "uploaded_file"
		outFile, err := os.Create(fileName)
		if err != nil {
			http.Error(w, "Unable to create file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		// Save the uploaded file content to the created file
		_, err = io.Copy(outFile, file)
		if err != nil {
			http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Log the file upload (for debugging or auditing purposes)
		fmt.Printf("File %s uploaded successfully\n", fileName)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
