package ports

import "Auth/internal/domain/user"

type TokenGenerator interface {
	GetToken(user *user.User)(string,error)
}