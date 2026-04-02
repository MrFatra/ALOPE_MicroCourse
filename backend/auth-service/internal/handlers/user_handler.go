package handlers

import (
	"alope-course/auth-service/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Get All List User without Pagination
// @Tags         users
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /users [get]
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
