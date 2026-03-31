package handlers

import (
	"alope-course/auth-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Email == "admin@alope.com" && input.Password == "password123" {
		token, _ := utils.GenerateToken(1)
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"token":  token,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"message": "Email atau password salah"})
}
