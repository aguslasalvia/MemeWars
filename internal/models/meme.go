package models

import "time"

type Meme struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	Text      string    `json:"text,omitempty"`
	CreatedOn time.Time `json:"created_on"`
}
