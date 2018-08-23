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

var UploadFile = func(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(0)

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
					Name: name,
					Url:  c.HostName + c.FilesNews + name,
				})
				go u.StoreFile(FilesNewsPath, name, file)
			}
		}

		go r.MultipartForm.RemoveAll()

		resp := u.Message(true, "Files uploaded successfully.")
		resp["files"] = filesList
		u.Respond(w, resp)
		return
	}

	resp := u.Message(false, "Failed to read data.")
	u.Respond(w, resp)
	return
}
