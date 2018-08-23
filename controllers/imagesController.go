package controllers

import (
	"fmt"
	"net/http"
	"sync"

	c "FITstorage/constants"
	"FITstorage/models"
	u "FITstorage/utils"
)

const ImageNewsPath = "public/images/news/"

var UploadImage = func(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(100000000000)

	var images []models.Image

	for key, files := range r.MultipartForm.File {
		wg := sync.WaitGroup{}

		if key == "news" {
			for index := range files {
				wg.Add(1)
				go func(index int) {

					name, n, err := u.StoreFile(ImageNewsPath, files[index])

					images = append(images, models.Image{
						Name: name,
						Url:  c.HostName + c.ImagesNews + name,
					})

					fmt.Println(images, index, name, n, err)

					wg.Done()
				}(index)
			}
		}

		wg.Wait()
	}

	resp := u.Message(true, "Image uploaded successfully.")
	resp["images"] = images
	u.Respond(w, resp)
}
