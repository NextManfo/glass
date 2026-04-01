package hw

import "os/exec"

func TakePhoto() []byte {
	cmd := exec.Command(
		"rpicam-still",
		"-o", "-",
		"--nopreview",
		"--timeout", "1000",
		"--width", "640",
		"--height", "480",
		"--quality", "50",
	)
	imgBytes, _ := cmd.Output()
	return imgBytes
}
