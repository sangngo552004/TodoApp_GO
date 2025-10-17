package dto_responses

type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}
