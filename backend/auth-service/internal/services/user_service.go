package services

import (
	"alope-course/auth-service/internal/models"
	"alope-course/auth-service/internal/repositories"
)

func GetAllUsers() ([]models.User, error) {

	users, err := repositories.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}
