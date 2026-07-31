package services

import (
	"memewars/internal/db"
	"memewars/internal/models"
)

type MemeService struct{}

// func (ms *MemeService) Create(meme models.Meme) error {

// }

func (ms *MemeService) FindMemeFromRoom(memeID uint, roomID uint) (*models.Meme, error) {
	var m models.Meme

	if err := db.DB.Where("id = ? AND room_id = ?", memeID, roomID).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (ms *MemeService) FindAllMemesFromRoom(roomID uint) []models.Meme {
	var mss []models.Meme
	if err := db.DB.Where("room_id = ?", roomID).Find(&mss).Error; err != nil {
		return []models.Meme{}
	}
	return mss
}
