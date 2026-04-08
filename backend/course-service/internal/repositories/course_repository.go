package repositories

import (
	"alope-course/course-service/internal/config"
	"alope-course/course-service/internal/models"
)

func GetCourseRepository() ([]models.Course, error) {
	db := config.DB

	var courses []models.Course

	err := db.Order("id ASC").Find(&courses).Error

	if err != nil {
		return nil, err
	}

	return courses, nil
}
