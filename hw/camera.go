package hw

import "os/exec"

func TakePhoto() []byte {
	cmd := exec.Command("rpicam-still", "-o", "-", "--nopreview", "--timeout", "1000")
	imgBytes, _ := cmd.Output()
	return imgBytes
}
