package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func saveAndRecordFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	timestamp := time.Now().Unix()
	uuidFilename := uuid.New().String()
	yearMonth := time.Now().Format("2006-01")

	filename := fmt.Sprintf("%s_%s_%d%s", uuidFilename, timestamp, filepath.Ext(file.Filename))

	uploadPath := filepath.Join("uploads", yearMonth)

	if err := os.MkdirAll(uploadPath, 0777); err != nil {
		return "", fmt.Errorf("failed to create upload folder: %v", err)
	}

	fullPath := filepath.Join(uploadPath, filename)

	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return filename, nil
}
