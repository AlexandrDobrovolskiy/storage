package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/chilts/sid"
)

func StoreFile(path string, file *multipart.FileHeader) (string, int64, error) {

	parseName := strings.Split(file.Filename, ".")
	ext := parseName[len(parseName)-1]

	name := sid.Id() + "." + ext

	f, err := os.Create(path + name)

	if err != nil {
		return "", 0, errors.New("ERROR WHILE CREATING NEW FILE")
	}

	image, err := file.Open()

	n, err := io.Copy(f, image)

	if err != nil {
		return "", 0, errors.New("ERROR WHILE SAVING IMAGE")
	}

	return name, n, nil
}
