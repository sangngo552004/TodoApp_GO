package handlers

import (
	"awesomeProject1/intelnal/dtos/dto_requests"
	"awesomeProject1/intelnal/dtos/dto_responses"
	"awesomeProject1/intelnal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (s *AuthHandler) Login(c *gin.Context) {
	var req dto_requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiError := &dto_responses.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		dto_responses.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", apiError)
		return
	}

	accessToken, refreshToken, user, err := s.service.Login(&req)
	if err != nil {
		c.Error(err)
		return
	}
	userResponse := dto_responses.UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}
	tokenResponse := dto_responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         userResponse,
	}
	dto_responses.SuccessResponse(c, http.StatusOK, "Login success", tokenResponse)
}

func (s *AuthHandler) Register(c *gin.Context) {
	var req dto_requests.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiError := &dto_responses.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		dto_responses.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", apiError)
		return
	}
	user, err := s.service.Register(&req)
	if err != nil {
		c.Error(err)
		return
	}
	userResponse := dto_responses.UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}
	dto_responses.SuccessResponse(c, http.StatusCreated, "Register success", userResponse)
}

func (s *AuthHandler) Refresh(c *gin.Context) {
	var req dto_requests.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiError := &dto_responses.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		dto_responses.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", apiError)
		return
	}
	accessToken, err := s.service.RefreshToken(&req)
	if err != nil {
		c.Error(err)
		return
	}
	refreshResponse := dto_responses.RefreshResponse{
		AccessToken: accessToken,
	}
	dto_responses.SuccessResponse(c, http.StatusOK, "Refresh success", refreshResponse)
}

func (s *AuthHandler) Logout(c *gin.Context) {
	var req dto_requests.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiError := &dto_responses.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		dto_responses.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", apiError)
		return
	}
	err := s.service.Logout(&req)
	if err != nil {
		c.Error(err)
		return
	}
	dto_responses.SuccessResponse(c, http.StatusOK, "Logout success", nil)
}
