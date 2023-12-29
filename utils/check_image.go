package utils

import (
	"fmt"
	"net/http"
	"os"
)

// CheckImage checks the image located at the specified path.
// It reads the file, detects the image type and size, and performs validations.
// The image size must be lower than 1MB, and the image type must be either JPEG or PNG.
// If the image fails any of the validations, an error is returned.
// Otherwise, it returns true indicating that the image is valid.
func CheckImage(path string) (bool, error) {

	bytes, err := os.ReadFile(path)

	if err != nil {
		return false, err
	}

	mimeType := http.DetectContentType(bytes)

	imageSize := len(bytes)

	//check image size. Must lower than 1MB

	if imageSize > 1000000 {
		return false, fmt.Errorf("image size is too large. Must lower than 1MB")
	}

	//check image type. Must be jpeg or png

	if mimeType != "image/jpeg" && mimeType != "image/png" {
		return false, fmt.Errorf("image type must be jpeg or png")
	}

	return true, nil

}
