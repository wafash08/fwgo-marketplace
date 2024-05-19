package helpers

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(file *multipart.FileHeader) string {
	// Tentukan lokasi penyimpanan file
	uploadDir := "src/uploads"
	epochTime := time.Now().Unix()
	filePath := filepath.Join(uploadDir, fmt.Sprintf("%d_%s", epochTime, file.Filename))

	// Buat direktori jika belum ada
	_, err := os.Stat(uploadDir)
	if os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}
	return filePath
}
