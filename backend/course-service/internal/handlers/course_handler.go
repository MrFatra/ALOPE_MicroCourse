package handlers

import (
	"alope-course/course-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourseHandler(c *gin.Context) {
	courses, err := services.GetCourseService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "ALP-003",
			"message": "Gagal mengambil data kursus.",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    "ALP-004",
		"message": "Data Kursus berhasil dimuat.",
		"data":    courses,
	})
}
