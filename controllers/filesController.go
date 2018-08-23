package controllers

import (
	"mime/multipart"
	"net/http"
	"sync"

	c "FITstorage/constants"
	"FITstorage/models"
	u "FITstorage/utils"
)

const FilesNewsPath = "public/files/news/"

var UploadFile = func(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(100000000000)

	for key, files := range r.MultipartForm.File {
		var filesList = make([]models.File, 0)
		if key == "news" {
			if len(files) == 0 {
				resp := u.Message(false, "No files given.")
				u.Respond(w, resp)
				return
			}
			wg := &sync.WaitGroup{}
			wg.Add(len(files))
			for _, file := range files {
				go func(wg *sync.WaitGroup, images *[]models.File, file *multipart.FileHeader) {
					name, _ := u.StoreFile(FilesNewsPath, file)
					*images = append(*images, models.File{
						Name: name,
						Url:  c.HostName + c.FilesNews + name,
					})
					wg.Done()
				}(wg, &filesList, file)
			}
			wg.Wait()
		}

		resp := u.Message(true, "Files uploaded successfully.")
		resp["files"] = filesList
		u.Respond(w, resp)
		return
	}

	resp := u.Message(false, "Failed to read data.")
	u.Respond(w, resp)
	return
}
