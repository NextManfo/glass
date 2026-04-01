package hw

import (
	"encoding/base64"
	"fmt"
	"os/exec"
)

type CameraPayload struct {
	Image string `json:"image"`
}

func TakePhoto() ([]byte, error) {
	cmd := exec.Command(
		"rpicam-still",
		"-o", "-",
		"--nopreview",
		"--timeout", "1000",
		"--width", "640",
		"--height", "480",
		"--quality", "50",
	)
	imgBytes, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("errore acquisizione foto: %w", err)
	}
	return imgBytes, nil
}

func TakePhotoBase64() (string, error) {
	imgBytes, err := TakePhoto()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imgBytes), nil
}

