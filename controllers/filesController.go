package controllers

import (
	"net/http"
	"strings"

	"github.com/google/uuid"

	c "FITstorage/constants"
	"FITstorage/models"
	u "FITstorage/utils"
)

const FilesNewsPath = "public/files/news/"

func UploadFile(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(0)

	if err != nil {
		resp := u.Message(false, "Failed to read data.")
		u.Respond(w, resp)
		return
	}

	for key, files := range r.MultipartForm.File {
		var filesList = make([]models.File, 0)
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
				filesList = append(filesList, models.File{
					OldName: file.Filename,
					Url:     c.HostName + c.FilesNews + name,
					Name:    name,
					Type:    ext,
				})
				go u.StoreFile(FilesNewsPath, name, file)
			}
		}

		resp := u.Message(true, "Files uploaded successfully.")
		resp["files"] = filesList
		u.Respond(w, resp)
		return
	}

	go r.MultipartForm.RemoveAll()

	resp := u.Message(false, "Failed to read data.")
	u.Respond(w, resp)
	return
}
