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

type RankingEntry struct {
	MemeID     uint   `json:"meme_id"`
	ImageURL   string `json:"image_url"`
	Text       string `json:"text"`
	TotalVotes int    `json:"total_votes"`
}
