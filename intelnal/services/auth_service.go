package services

import (
	"awesomeProject1/intelnal/apperror"
	"awesomeProject1/intelnal/dtos/dto_requests"
	"awesomeProject1/intelnal/models"
	"awesomeProject1/intelnal/repositories"
	"awesomeProject1/intelnal/utils"
	"context"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req *dto_requests.RegisterRequest) (*models.User, error)
	Login(req *dto_requests.LoginRequest) (string, string, *models.User, error)
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

func (s *AuthServiceImpl) Register(req *dto_requests.RegisterRequest) (*models.User, error) {
	_, err := s.userRepository.FindByEmail(req.Email)
	if err == nil {
		return nil, apperror.Conflict("Email already exists", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperror.InternalServerError("Error hashing password", err)
	}
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	if err := s.userRepository.Create(&user); err != nil {
		return nil, apperror.InternalServerError("Error creating user", err)
	}
	return &user, nil
}

func (s *AuthServiceImpl) Login(req *dto_requests.LoginRequest) (string, string, *models.User, error) {
	user, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {
		return "", "", nil, apperror.NotFound("Email not found", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", nil, apperror.Unauthorized("Invalid credentials", err)
	}
	access, _ := utils.GenerateAccessToken(user.ID)
	refresh, _ := utils.GenerateRefreshToken(user.ID, s.redisClient)

	return access, refresh, user, nil
}

func (s *AuthServiceImpl) RefreshToken(req *dto_requests.RefreshRequest) (string, error) {
	claims, err := utils.ValidateRefreshToken(req.RefreshToken, s.redisClient)
	if err != nil {
		return "", apperror.Unauthorized("Invalid refresh token", err)
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return "", apperror.InternalServerError("Invalid token payload", err)
	}

	newAccess, err := utils.GenerateAccessToken(uint(userID))
	if err != nil {
		return "", apperror.InternalServerError("Error generating new access token", err)
	}

	return newAccess, nil
}

func (s *AuthServiceImpl) Logout(req *dto_requests.RefreshRequest) error {
	ctx := context.Background()
	if err := s.redisClient.Del(ctx, req.RefreshToken).Err(); err != nil {
		return apperror.InternalServerError("Error deleting refresh token", err)
	}
	return nil
}
