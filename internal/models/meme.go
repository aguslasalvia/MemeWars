package models

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type Meme struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	RoomID uint `json:"room_id"`
	Room   Room `json:"room" gorm:"foreignKey:RoomID"`

	ImageURL string `json:"image_url"`
	Text     string `json:"text"`

	Votes []Vote `json:"votes,omitempty" gorm:"foreignKey:MemeID"`
}

type CreateMemeRequest struct {
	UserID uint                  `form:"user_id" binding:"required"`
	Text   string                `form:"text"`
	Image  *multipart.FileHeader `form:"image" binding:"required"`
}
