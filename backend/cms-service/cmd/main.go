package main

import (
	"alope-course/cms-service/internal/config"
	"alope-course/cms-service/internal/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.SetupRouter(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	r.Run(":" + port)
}
