package security

import (
	"Auth/internal/domain/user"
	
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JwtAdapter struct{
	secretKey string
}


func NewJwtAdapter(key string) *JwtAdapter{
	return &JwtAdapter{secretKey: key,}
}

func (j *JwtAdapter) GetToken(user *user.User)(string,error){
	claims:= jwt.MapClaims{
		"sub": user.GetId(),
		"email": user.Email(),
		"role": user.Role(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
} 