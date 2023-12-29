package utils

import (
	"fmt"
	"net/http"
	"os"
)

func CheckImage(path string) (bool, error) {

	bytes, err := os.ReadFile(path)

	if err != nil {
		return false, err
	}

	mimeType := http.DetectContentType(bytes)

	imageSize := len(bytes)

	//check image size. Must lower than 1MB

	if imageSize > 1000000 {
		return false, fmt.Errorf("Image size is too large. Must lower than 1MB")
	}

	//check image type. Must be jpeg or png

	if mimeType != "image/jpeg" && mimeType != "image/png" {
		return false, fmt.Errorf("Image type must be jpeg or png")
	}

	return true, nil

}
