package controllers

import (
	c "FITstorage/constants"
	"FITstorage/models"
	u "FITstorage/utils"
	"mime/multipart"
	"net/http"
	"sync"
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
			wg := &sync.WaitGroup{}
			wg.Add(len(files))
			for _, file := range files {
				go func(wg *sync.WaitGroup, images *[]models.Image, file *multipart.FileHeader) {
					name, _ := u.StoreFile(ImageNewsPath, file)
					*images = append(*images, models.Image{
						Name: name,
						Url:  c.HostName + c.ImagesNews + name,
					})
					wg.Done()
				}(wg, &imagesList, file)
			}
			wg.Wait()

			resp := u.Message(true, "Image uploaded successfully.")
			resp["images"] = imagesList
			u.Respond(w, resp)
			return
		}
	}

	resp := u.Message(false, "Failed to read data.")
	u.Respond(w, resp)
	return
}
