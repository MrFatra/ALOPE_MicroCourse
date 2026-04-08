package models

import (
	"time"

	"gorm.io/gorm"
)

type Testimonial struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CourseID  uint           `json:"course_id"`
	Course    Course         `gorm:"foreignKey:CourseID" json:"course"`
	UserID    uint           `json:"user_id"`
	User      User           `gorm:"foreignKey:user_id" json:"user"`
	Message   string         `gorm:"type:text" json:"message"`
	Rating    uint8          `gorm:"type:smallint" json:"rating"`
	Status    string         `gorm:"type:testimonial_status;default:'pending'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
