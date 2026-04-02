package routes

import (
	"alope-course/auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/users", handlers.GetAllUsers)

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", handlers.Login)
	}

	return r
}
