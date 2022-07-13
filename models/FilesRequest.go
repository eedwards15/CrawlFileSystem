package models

import (
	"time"
)

type FilesRequest struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdOn"`
	Extension   string    `json:"extension"`
}

func NewFileRequest(name string, fileLocation string, ext string) *FilesRequest {
	var request = &FilesRequest{
		Name:        name,
		Location:    fileLocation,
		Description: "",
		Extension:   ext,
		CreatedOn:   time.Now(),
	}

	return request
}
