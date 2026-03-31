package routes

import (
	"alope-course/course-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", controllers.GetPing)
	}

	return r
}
