package handlers

import (
	"alope-course/auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"code":    "ALP-002",
			"message": "Gagal mengambil data user",
			"data":    err.Error(), // nanti gak boleh kasih detail error ke user ya ham
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"code":    "ALP-001",
		"message": "Berhasil mengambil data user",
		"data":    users,
	})
}
