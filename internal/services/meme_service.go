package services

import (
	"errors"
	"fmt"
	"memewars/internal/db"
	"memewars/internal/models"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type MemeService struct{}

func (ms *MemeService) CreateMeme(roomID uint, userID uint, text string, file *multipart.FileHeader) (*models.Meme, error) {
	var room models.Room
	if err := db.DB.Where("active = ?", true).First(&room, roomID).Error; err != nil {
		return nil, errors.New("room not found or not active")
	}

	dir := fmt.Sprintf("static/memes/room_%d", roomID)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}

	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	fullPath := filepath.Join(dir, filename)

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return nil, err
	}

	meme := models.Meme{
		RoomID:   roomID,
		UserID:   userID,
		ImageURL: "/" + fullPath,
		Text:     text,
	}

	if err := db.DB.Create(&meme).Error; err != nil {
		return nil, err
	}

	return &meme, nil
}

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
