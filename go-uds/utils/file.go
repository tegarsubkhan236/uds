package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveFile(fileHeader *multipart.FileHeader) (string, error) {
	dst := "/home/fauzi/www/github/uds/upload"
	relativePath := filepath.Join(dst, fileHeader.Filename)

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	outFile, err := os.Create(relativePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return "./" + relativePath, nil
}

func RemoveFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}
