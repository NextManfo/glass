package hw

import "os/exec"

func TakePhoto() []byte {
	cmd := exec.Command("libcamera-still", "-o", "-", "--nopreview")
	imgBytes, _ := cmd.Output()
	return imgBytes
}
