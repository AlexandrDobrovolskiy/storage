package models

import (
	"github.com/jinzhu/gorm"
		 u "storage/utils"
)

type Image struct {
	gorm.Model
	Name string `json:"name"`
	Url string `json:"url"`
}

func (image *Image) Store() (map[string] interface{}) {

	resp := u.Message(true, "Image successfully stored")
	resp["image"] = image

	return resp
}