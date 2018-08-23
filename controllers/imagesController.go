package controllers

import (
	"net/http"
	"strings"

	"github.com/google/uuid"

	c "FITstorage/constants"
	"FITstorage/models"
	u "FITstorage/utils"
)

const ImageNewsPath = "public/images/news/"

func UploadImage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(0)

	if err != nil {
		resp := u.Message(false, "Failed to read data.")
		u.Respond(w, resp)
		return
	}

	for key, files := range r.MultipartForm.File {
		var imagesList = make([]models.Image, 0)
		if key == "news" {
			if len(files) == 0 {
				resp := u.Message(false, "No files given.")
				u.Respond(w, resp)
				return
			}

			for _, file := range files {
				parseName := strings.Split(file.Filename, ".")
				ext := parseName[len(parseName)-1]
				name := uuid.New().String() + "." + ext
				imagesList = append(imagesList, models.Image{
					Name: name,
					Url:  c.HostName + c.FilesNews + name,
				})
				go u.StoreFile(ImageNewsPath, name, file)
			}

			resp := u.Message(true, "Image uploaded successfully.")
			resp["images"] = imagesList
			u.Respond(w, resp)
			return
		}
	}

	go r.MultipartForm.RemoveAll()

	resp := u.Message(false, "Failed to read data.")
	u.Respond(w, resp)
	return
}
