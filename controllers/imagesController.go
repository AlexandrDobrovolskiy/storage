package controllers

import (
	"net/http"
	u "storage/utils"
	"fmt"
)

const path = "public/images/news"

var UploadImage = func( w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(100000000000)

	for key, files := range r.MultipartForm.File{
		if key == "news" {
			for index := range files {
				go func() {
					n, e := u.StoreFile(path, files[index])

					fmt.Println(index, n, e)
				}()
			}
		}
	}

	resp := u.Message(true, "Image uploaded successfully.")
	u.Respond(w, resp)
}
