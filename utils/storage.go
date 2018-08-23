package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/google/uuid"
)

func StoreFile(path string, file *multipart.FileHeader) (string, error) {

	parseName := strings.Split(file.Filename, ".")
	ext := parseName[len(parseName)-1]

	name := uuid.New().String() + "." + ext

	f, err := os.Create(path + name)

	if err != nil {
		return "", errors.New("failed to create file")
	}

	defer f.Close()

	image, err := file.Open()

	if err != nil {
		return "", errors.New("failed to read file")
	}

	defer image.Close()

	_, err = io.Copy(f, image)

	if err != nil {
		return "", errors.New("failed to copy file")
	}

	return name, nil
}
