package main

import "time"

type FilesRequest struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdOn"`
	Extension   string    `json:"extension"`
}
