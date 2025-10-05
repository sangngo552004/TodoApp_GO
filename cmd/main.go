package main

import (
	"awesomeProject1/intelnal/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// Kết nối DB
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to MySQL database")

}
