package services

import (
	"awesomeProject1/intelnal/DTOs/DTOrequest"
	"awesomeProject1/intelnal/config"
	"awesomeProject1/intelnal/models"
	"awesomeProject1/intelnal/repositories"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req *DTOrequest.RegisterRequest) error
	Login(req *DTOrequest.LoginRequest) (string, error)
}

type AuthServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &AuthServiceImpl{userRepository: userRepository}
}

func (s *AuthServiceImpl) Register(req *DTOrequest.RegisterRequest) error {
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

func (s *AuthServiceImpl) Login(req *DTOrequest.LoginRequest) (string, error) {
	user, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("Email not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("Invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(config.TokenExpireDuration).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET", "secret")))
	if err != nil {
		return "", errors.New("Error signing token")
	}
	return tokenString, nil
}
