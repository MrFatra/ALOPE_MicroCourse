package handlers

import (
	"alope-course/course-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTestimonialHandler(c *gin.Context) {
	testimonials, err := services.GetTestimonialService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "ALP-004",
			"message": "Gagal mengambil data testimonial.",
			"data":    err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    "ALP-006",
		"message": "Berhasil memuat data testimonial.",
		"data":    testimonials,
	})

}
