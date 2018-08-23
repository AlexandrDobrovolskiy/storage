package controllers

import (
	"net/http"
	u "storage/utils"
	"fmt"
	"storage/models"
	"sync"
	c "storage/constants"
)

const FilesNewsPath = "public/files/news/"

var UploadFile = func(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(100000000000)

	var fResp []models.File

	for key, files := range r.MultipartForm.File{
		wg := sync.WaitGroup{}

		if key == "news" {
			for index := range files {
				wg.Add(1)
				go func(index int) {

					name, n, err := u.StoreFile(FilesNewsPath, files[index])

					fResp = append(fResp, models.File{
						Name: name,
						Url: c.HostName + c.FilesNews + name,
					})

					fmt.Println(files, index, name, n, err)

					wg.Done()
				}(index)
			}
		}

		wg.Wait()
	}

	resp := u.Message(true, "File uploaded successfully.")
	resp["images"] = fResp
	u.Respond(w, resp)
}
