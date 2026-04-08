package handlers

import (
	"alope-course/course-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModuleHandler(c *gin.Context) {
	modules, err := services.GetModuleService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "ALP-004",
			"message": "Gagal mengambil data modul.",
			"data":    err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    "ALP-005",
		"message": "Berhasil memuat data modul.",
		"data":    modules,
	})
}
