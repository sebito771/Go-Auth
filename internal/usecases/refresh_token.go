package usecases

import (
	"Auth/internal/ports"
	"errors"
	"time"
)

type RefreshTokenUseCase struct {
	repo      ports.UserRepository
	tokenGen  ports.TokenGenerator
	tokenVal  ports.TokenValidator
	blackList ports.TokenBlackList
}

var (
	ErrInvalidToken     = errors.New("invalid or expired token")
	ErrTokenInBlackList = errors.New("token has been invalidated")
)

func NewRefreshTokenUseCase(repo ports.UserRepository, tokenGen ports.TokenGenerator, tokenVal ports.TokenValidator, blackList ports.TokenBlackList) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		repo:      repo,
		tokenGen:  tokenGen,
		tokenVal:  tokenVal,
		blackList: blackList,
	}
}

func (r *RefreshTokenUseCase) Refresh(token string) (string, error) {
	// Check if token is in blacklist (user logged out)
	isBlackListed, err := r.blackList.IsBlackListed(token)
	if err != nil {
		return "", err
	}
	if isBlackListed {
		return "", ErrTokenInBlackList
	}

	// Validate token
	claims, err := r.tokenVal.ValidateToken(token)
	if err != nil {
		return "", ErrInvalidToken
	}

	// Check if token is expired
	if claims.Exp < time.Now().Unix() {
		return "", ErrInvalidToken
	}

	// Get user from database
	user, err := r.repo.FindById(claims.UserdID)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Generate new token
	newToken, err := r.tokenGen.GetToken(user)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
