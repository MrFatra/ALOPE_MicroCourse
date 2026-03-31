package main

import (
	"alope-course/auth-service/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Info: .env file tidak ditemukan, menggunakan env system")
	}

	// Set Mode
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	//
	log.Printf("[AUTH-SERVICE] Server berjalan di port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
