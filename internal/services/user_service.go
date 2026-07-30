package services

import (
	"memewars/internal/db"
	"memewars/internal/models"
)

type UserService struct{}

func (us *UserService) Create(u *models.User) (*models.User, error) {
	user := models.User{
		Name: u.Name,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GetUserByName(name string) (*models.User, error) {
	var u models.User
	if err := db.DB.First(&u, name).Error; err != nil {
		return nil, err
	}

	return &u, nil
}
