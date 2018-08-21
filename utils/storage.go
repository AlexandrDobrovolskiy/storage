package utils

import (
	"os"
	"io"
	"mime/multipart"
	"errors"
	)

func StoreFile(path string, file *multipart.FileHeader) (int64, error) {
	f, err := os.Create(path + file.Filename)

	if err != nil {
		return 0, errors.New("ERROR WHILE CREATING NEW FILE")
	}

	image, err := file.Open()

	n, err := io.Copy(f, image)

	if err != nil {
		return 0, errors.New("ERROR WHILE SAVING IMAGE")
	}

	return n, nil
}

