package services

import (
	"memewars/internal/db"
	"memewars/internal/models"
)

type MemeService struct{}

func (ms *MemeService) FindMemeFromRoom(memeID uint, roomID uint) (*models.Meme, error) {
	var m models.Meme

	if err := db.DB.Where("id = ? AND room_id = ?", memeID, roomID).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}
