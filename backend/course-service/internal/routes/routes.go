package routes

import (
	"alope-course/course-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/courses", handlers.GetCourseHandler)
		api.GET("/modules", handlers.GetModuleHandler)
		api.GET("/testimonials", handlers.GetTestimonialHandler)
	}

	return r
}
