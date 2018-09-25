package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/uuid"

	c "FITstorage/constants"
	u "FITstorage/utils"
)

const FilePondTempPath = "public/files/temp/"
const FilePondPath = "public/files/news/"

func Options(w http.ResponseWriter, r *http.Request) {
	u.RespondText(w, "pong")
}

func FilePondProcess(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(0)

	if err != nil {
		u.RespondText(w, "No files given")
		return
	}

	for _, files := range r.MultipartForm.File {

		if len(files) == 0 {
			return
		}

		name := uuid.New().String()
		for _, file := range files {
			dir := FilePondTempPath + name

			os.Mkdir(dir, os.ModePerm)
			go u.StoreFile(dir+"/", file.Filename, file)
		}

		u.RespondText(w, name)
		return
	}

	go r.MultipartForm.RemoveAll()
	u.RespondText(w, "Failed")
	return
}

func FilePondDelete(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		u.RespondText(w, "Delete failed")
		return
	}

	toDelete := string(body)

	go os.RemoveAll(FilePondTempPath + toDelete)

	u.RespondText(w, "Deleted")
	return
}

type ConfirmMessage struct {
	Files []string `json:"files"`
}

func SubmitStore(w http.ResponseWriter, r *http.Request) {
	req := &ConfirmMessage{}
	var filesList = make([]string, 0)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		u.Respond(w, u.Message(false, "Bad request."))
		return
	}

	for _, id := range req.Files {
		dest := FilePondPath + string(id)
		os.Mkdir(dest, os.ModePerm)

		src := FilePondTempPath + string(id)

		err := u.CopyDir(src, dest)

		directory, _ := os.Open(dest)
		files, err := directory.Readdir(-1)

		if err != nil {
			u.Respond(w, u.Message(false, "Something went wrong."))
			return
		}

		filesList = append(filesList, c.HostName+c.FilesNews+id+"/"+files[0].Name())

		if err != nil {
			u.Respond(w, u.Message(false, "Error while copying directory."))
			return
		}

		go os.RemoveAll(src)
	}

	resp := u.Message(true, "Confirmation successfull.")
	resp["files"] = filesList

	u.Respond(w, resp)
}
