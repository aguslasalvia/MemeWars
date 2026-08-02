package services

import (
	"errors"
	"memewars/internal/db"
	"memewars/internal/models"
)

type RoomService struct{}

func (rs *RoomService) CreateRoom(name string) (*models.Room, error) {
	if name == "" {
		return nil, errors.New("Room name is empty")
	}

	room := models.Room{
		Name:   name,
		Active: true,
	}

	if err := db.DB.Create(&room).Error; err != nil {
		return nil, err
	}

	return &room, nil

}

func (rs *RoomService) GetRoom(room_id int) (*models.Room, error) {
	var r models.Room
	if err := db.DB.Preload("Memes").Where("active = ?", true).First(&r, room_id).Error; err != nil {
		return nil, err
	}
	return &r, nil

}

func (rs *RoomService) UploadMeme() {

}

func (rs *RoomService) GetRanking(roomID uint) ([]models.RankingEntry, error) {
	results := []models.RankingEntry{}

	err := db.DB.
		Table("memes").
		Select("memes.id as meme_id, memes.image_url, memes.text, COALESCE(SUM(votes.value), 0) as total_votes").
		Joins("LEFT JOIN votes ON votes.meme_id = memes.id").
		Where("memes.room_id = ?", roomID).
		Group("memes.id").
		Order("total_votes desc").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
