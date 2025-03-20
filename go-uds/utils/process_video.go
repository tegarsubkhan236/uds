package utils

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

const (
	videoDir               = "/home/fauzi/www/github/uds/upload/videos/"
	maxVideoSize           = 2 << 30 // 2GB
	allowedVideoExtensions = ".mp4"
)

// Mp4ToFMp4 converts an MP4 file to HLS with fMP4 segments using ffmpeg.
func Mp4ToFMp4(fileHeader *multipart.FileHeader) (string, error) {
	err := validateVideo(fileHeader)
	if err != nil {
		return "", err
	}

	hlsDir := filepath.Join(videoDir, "hls_"+strconv.Itoa(int(time.Now().Unix())))
	err = os.MkdirAll(hlsDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	initFileName := "init.mp4"
	playlistPath := filepath.Join(hlsDir, "playlist.m3u8")

	srcFile, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	originalFilePath := filepath.Join(hlsDir, fileHeader.Filename)
	outFile, err := os.Create(originalFilePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, srcFile)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("ffmpeg",
		"-i", originalFilePath,
		"-c:v", "copy", // libx264 | copy
		"-c:a", "copy", // cca | copy
		"-hls_time", "10",
		"-hls_playlist_type", "vod",
		"-hls_segment_type", "fmp4",
		"-hls_fmp4_init_filename", initFileName,
		"-f", "hls",
		playlistPath,
	)

	logFile, err := os.Create(filepath.Join(hlsDir, "ffmpeg_output.log"))
	if err != nil {
		return "", err
	}
	defer logFile.Close()

	cmd.Stdout = logFile
	cmd.Stderr = logFile

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	cmd.Start()      // Start the process
	err = cmd.Wait() // Wait for the process to finish
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("video conversion timed out")
		}
		return "", fmt.Errorf("failed to convert video to HLS: %v", err)
	}

	err = os.Remove(originalFilePath)
	if err != nil {
		fmt.Printf("Warning: Failed to delete original file: %v\n", err)
	}

	return hlsDir, nil
}

func validateVideo(header *multipart.FileHeader) error {
	size := header.Size
	if size > maxVideoSize {
		return errors.New("video too large")
	}

	extension := filepath.Ext(header.Filename)
	if extension != allowedVideoExtensions {
		return errors.New("file type not allowed")
	}

	return nil
}
