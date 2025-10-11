package utils

import (
	"awesomeProject1/intelnal/config"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var accessSecret = []byte(config.GetEnv("JWT_SECRET", "secret"))
var refreshSecret = []byte(config.GetEnv("JWT_REFRESH_SECRET", "refresh"))

var accessExpireDuration = time.Minute * 5
var refreshExpireDuration = time.Hour * 24

func GenerateAccessToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     accessExpireDuration,
	})
	tokenString, err := token.SignedString(accessSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     refreshExpireDuration,
	})
	tokenString, err := token.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ValidateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	return ParseToken(tokenString, accessSecret)
}
func ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	return ParseToken(tokenString, refreshSecret)
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
