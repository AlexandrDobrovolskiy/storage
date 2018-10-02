package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/uuid"

	c "FITstorage/constants"
	"FITstorage/models"
	u "FITstorage/utils"
)

const FilesNewsPath = "public/files/news/"

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(0)

	if err != nil {
		resp := u.Message(false, "Failed to read data.")
		u.Respond(w, resp)
		return
	}

	println(formatRequest(r))

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

				parsed, _ := url.Parse(file.Filename)

				filename := parsed.Path

				filesList = append(filesList, models.File{
					Url: c.HostName + c.FilesNews + dirName + "/" + string(filename),
				})

				os.Mkdir(dir, os.ModePerm)
				go u.StoreFile(dir+"/", string(filename), file)
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
