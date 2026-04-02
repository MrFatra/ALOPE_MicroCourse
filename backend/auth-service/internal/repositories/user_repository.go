package repositories

import (
	"alope-course/auth-service/internal/config"
	model "alope-course/auth-service/internal/models"
)

func GetUsers() ([]model.User, error) {
	db := config.DB
	var users []model.User

	err := db.Order("id DESC").Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
