package models

type File struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	OldName string `json:"oldName"`
	Type    string `json:"type"`
}
