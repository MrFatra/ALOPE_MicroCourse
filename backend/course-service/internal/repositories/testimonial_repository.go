package repositories

import (
	"alope-course/course-service/internal/config"
	"alope-course/course-service/internal/models"
)

func GetTestimonialRepository() ([]models.Testimonial, error) {
	db := config.DB

	var testimonials []models.Testimonial

	err := db.
		Preload("Course").
		Preload("Course.Category").
		Find(&testimonials).Error

	if err != nil {
		return nil, err
	}

	return testimonials, nil
}
