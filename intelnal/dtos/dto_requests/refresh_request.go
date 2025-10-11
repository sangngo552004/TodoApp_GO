package dto_requests

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
