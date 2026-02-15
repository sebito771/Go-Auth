package ports

import "Auth/internal/domain/user"

type TokenGenerator interface {
	GetToken(user *user.User)(string,error)
}

//token claims

type TokenClaims struct{
	UserdID int64
	Email string
	Role string
	Exp int64
}


type TokenValidator interface{
	ValidateToken(tokenstr string)(*TokenClaims,error)
}