package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

func StoreFile(path string, name string, file *multipart.FileHeader) error {

	f, err := os.Create(path + name)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer f.Close()

	image, err := file.Open()

	if err != nil {
		return errors.New("failed to read file")
	}

	defer image.Close()

	_, err = io.Copy(f, image)

	if err != nil {
		return errors.New("failed to copy file")
	}

	return nil
}
