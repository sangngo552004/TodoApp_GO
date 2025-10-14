package main

import (
	"awesomeProject1/intelnal/config"
	"awesomeProject1/intelnal/handlers"
	"awesomeProject1/intelnal/repositories"
	"awesomeProject1/intelnal/routes"
	"awesomeProject1/intelnal/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// Kết nối DB
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to MySQL database")

	redisClient, err := config.ConnectRedis()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}
	fmt.Println("Connected to Redis")

	r := gin.Default()
	// Khởi tạo repository, service, handler
	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)
	routes.InitTodoRoutes(r, todoHandler)

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, redisClient)
	authHandler := handlers.NewAuthHandler(authService)
	routes.InitAuthRoutes(r, authHandler)

	port := config.GetEnv("SERVER_PORT", "8080")
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}
