package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	imageDir               = "./storage/public/images/"
	maxImageSize           = 25 << 20 // 25MB
	allowedImageExtensions = ".mp4"
)

func SaveAndConvertToWebP(fileHeader *multipart.FileHeader) (string, error) {
	err := validateImage(fileHeader)
	if err != nil {
		return "", err
	}

	dstDir := imageDir

	originalFileName := fileHeader.Filename
	originalFilePath := filepath.Join(dstDir, originalFileName)

	srcFile, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	outFile, err := os.Create(originalFilePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, srcFile)
	if err != nil {
		return "", err
	}

	webpFileName := fmt.Sprintf("%s.webp", originalFileName)
	webpFilePath := filepath.Join(dstDir, webpFileName)

	cmd := exec.Command("ffmpeg",
		"-i", originalFilePath,
		"-c:v", "libwebp",
		"-quality", "80",
		"-lossless", "0",
		webpFilePath,
	)

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("gagal mengonversi ke WebP: %v", err)
	}

	err = os.Remove(originalFilePath)
	if err != nil {
		return "", fmt.Errorf("gagal menghapus file asli: %v", err)
	}

	return "./" + webpFilePath, nil
}

func validateImage(header *multipart.FileHeader) error {
	size := header.Size
	if size > maxVideoSize {
		return errors.New("image too large")
	}

	extension := filepath.Ext(header.Filename)
	if extension != allowedVideoExtensions {
		return errors.New("file type not allowed")
	}

	return nil
}
