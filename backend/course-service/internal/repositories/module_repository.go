package repositories

import (
	"alope-course/course-service/internal/config"
	"alope-course/course-service/internal/models"
)

func GetModuleRepository() ([]models.Module, error) {
	db := config.DB

	var modules []models.Module

	err := db.Order("id ASC").Find(&modules).Error

	if err != nil {
		return nil, err
	}

	return modules, nil
}
