package controllers

import (
	"net/http"
	"os"

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
				dirName := uuid.New().String()

				dir := FilesNewsPath + dirName

				filesList = append(filesList, models.File{
					Url: c.HostName + c.FilesNews + dirName + "/" + file.Filename,
				})

				os.Mkdir(dir, os.ModePerm)
				go u.StoreFile(dir+"/", file.Filename, file)
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
