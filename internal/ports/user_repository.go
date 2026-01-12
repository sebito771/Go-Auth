package ports

import "Auth/internal/domain/user"

type UserRepository interface {
	Save(user *user.User) error
	FindByEmail(email string) (*user.User, error)
}


