package models

import "time"

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	State     string    `json:"state"`
	CreatedOn time.Time `json:"created_on"`
}

type CreateRoomRequest struct {
	Name string `json:"name" binding:"required"`
}
