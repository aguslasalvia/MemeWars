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

func (rs *RoomService) GetRanking() {

}
