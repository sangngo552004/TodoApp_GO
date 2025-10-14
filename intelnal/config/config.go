package config

import (
	"awesomeProject1/intelnal/models"
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetEnv(key, defaulValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaulValue
	}
	return value
}

const TokenExpireDuration = time.Minute * 6

func ConnectDB() (*gorm.DB, error) {
	dsn := GetEnv("DB_USER", "root") + ":" +
		GetEnv("DB_PASSWORD", "root") + "@tcp(" +
		GetEnv("DB_HOST", "localhost") + ":" +
		GetEnv("DB_PORT", "3306") + ")/" +
		GetEnv("DB_NAME", "todo_db") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectRedis() (*redis.Client, error) {
	addr := GetEnv("REDIS_ADDR", "localhost:6379")
	password := GetEnv("REDIS_PASSWORD", "")
	db := 0

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
