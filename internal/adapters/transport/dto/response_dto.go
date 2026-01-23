package dto

import "Auth/internal/domain/user"

type UserResponse struct {
	Email   string `json:"email"`
	Role    user.Role `json:"role"`
	Message string `json:"message"`
}