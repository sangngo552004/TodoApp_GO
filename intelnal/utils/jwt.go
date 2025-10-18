package utils

import (
	"awesomeProject1/intelnal/config"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

var accessSecret = []byte(config.GetEnv("JWT_SECRET", "secret"))
var refreshSecret = []byte(config.GetEnv("JWT_REFRESH_SECRET", "refresh"))

var accessExpireDuration = time.Minute * 5
var refreshExpireDuration = time.Hour * 24

func GenerateAccessToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(accessExpireDuration).Unix(),
	})
	tokenString, err := token.SignedString(accessSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(userId uint, redisClient *redis.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(refreshExpireDuration).Unix(),
	})
	tokenString, err := token.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}
	err = redisClient.Set(context.Background(), tokenString, userId, refreshExpireDuration).Err()
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ValidateRefreshToken(tokenString string, redisClient *redis.Client) (jwt.MapClaims, error) {
	_, err := redisClient.Get(context.Background(), tokenString).Uint64()
	if err != nil {
		return nil, err
	}
	return ParseToken(tokenString, refreshSecret)
}
func ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	return ParseToken(tokenString, accessSecret)
}

func ParseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
