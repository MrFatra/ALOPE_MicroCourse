package seeders

import (
	model "alope-course/course-service/internal/models"
	"log"

	helpers "alope-course/course-service/internal/utils"

	"gorm.io/gorm"
)

func SeedCourses(db *gorm.DB) error {
	courses := []model.Course{
		{
			Title:       "LWD PBK",
			Description: "Belajar Web Programming dengan PBK.",
		},
		{
			Title:       "HTML Dasar",
			Description: "Belajar HTML dari dasar.",
		},
		{
			Title:       "CSS Dasar",
			Description: "Dasar-dasar CSS untuk pemula.",
		},
		{
			Title:       "CSS Layouting",
			Description: "Latihan Membuat Website Statis dengan HTML & CSS.",
		},
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for i := range courses {
			courses[i].Slug = helpers.Slugify(courses[i].Title)
			if err := tx.
				Where(model.Course{Title: courses[i].Title}).
				FirstOrCreate(&courses[i]).Error; err != nil {
				return err
			}
		}

		log.Println("Seeding Courses: Success")
		return nil
	})
}
