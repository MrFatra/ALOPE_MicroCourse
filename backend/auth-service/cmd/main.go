package main

import (
	"alope-course/auth-service/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load environment variables dari file .env
	// Jika menggunakan Docker, variabel ini biasanya dioper via docker-compose
	err := godotenv.Load()
	if err != nil {
		log.Println("Info: .env file tidak ditemukan, menggunakan env system")
	}

	// 2. Set mode Gin (release atau debug)
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 3. Inisialisasi Router dari package internal/routes
	r := routes.SetupRouter()

	// 4. Pengaturan Port
	// Auth service kita set default di 8081 agar tidak bentrok dengan Course service (8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// 5. Jalankan Server
	log.Printf("[AUTH-SERVICE] Server berjalan di port %s", port)
	
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}