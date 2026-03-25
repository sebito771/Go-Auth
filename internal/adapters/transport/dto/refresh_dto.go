package dto

type RefreshRequest struct {
	Token string `json:"token" binding:"required"`
}

type RefreshResponse struct {
	Token string `json:"token"`
}