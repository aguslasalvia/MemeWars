package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	MemeID uint `json:"meme_id"`
	Meme   Meme `json:"-" gorm:"foreignKey:MemeID"`

	UserID uint `json:"user_id"`
	User   User `json:"-" gorm:"foreignKey:UserID"`

	Value int `json:"value"`
}
