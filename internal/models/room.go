package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Memes  []Meme `json:"memes,omitempty" gorm:"foreignKey:RoomID"`
}

type CreateRoomRequest struct {
	Name string `json:"name" binding:"required"`
}
