package services

import (
	"awesomeProject1/intelnal/dtos/dto_requests"
	"awesomeProject1/intelnal/models"
	"awesomeProject1/intelnal/repositories"
	"awesomeProject1/intelnal/utils"
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req *dto_requests.RegisterRequest) error
	Login(req *dto_requests.LoginRequest) (string, string, error)
	RefreshToken(req *dto_requests.RefreshRequest) (string, error)
	Logout(req *dto_requests.RefreshRequest) error
}

type AuthServiceImpl struct {
	userRepository repositories.UserRepository
	redisClient    *redis.Client
}

func NewAuthService(userRepository repositories.UserRepository, redisClient *redis.Client) AuthService {
	return &AuthServiceImpl{
		userRepository: userRepository,
		redisClient:    redisClient,
	}
}

func (s *AuthServiceImpl) Register(req *dto_requests.RegisterRequest) error {
	_, err := s.userRepository.FindByEmail(req.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	if err := s.userRepository.Create(&user); err != nil {
		return errors.New("error creating user")
	}
	return nil
}

func (s *AuthServiceImpl) Login(req *dto_requests.LoginRequest) (string, string, error) {
	user, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {
		return "", "", errors.New("Email not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", errors.New("Invalid password")
	}
	access, _ := utils.GenerateAccessToken(user.ID)
	refresh, _ := utils.GenerateRefreshToken(user.ID, s.redisClient)

	return access, refresh, nil
}

func (s *AuthServiceImpl) RefreshToken(req *dto_requests.RefreshRequest) (string, error) {
	claims, err := utils.ValidateRefreshToken(req.RefreshToken, s.redisClient)
	if err != nil {
		return "", errors.New("invalid or expired refresh token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return "", errors.New("invalid token payload")
	}

	newAccess, err := utils.GenerateAccessToken(uint(userID))
	if err != nil {
		return "", errors.New("failed to generate new token")
	}

	return newAccess, nil
}

func (s *AuthServiceImpl) Logout(req *dto_requests.RefreshRequest) error {
	ctx := context.Background()
	if err := s.redisClient.Del(ctx, req.RefreshToken).Err(); err != nil {
		return err
	}
	return nil
}
