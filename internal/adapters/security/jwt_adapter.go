package security

import (
	"Auth/internal/domain/user"
	"Auth/internal/ports"
	"errors"

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

func (j *JwtAdapter) ValidateToken(tokenstr string)(*ports.TokenClaims,error){
  
	token,err:= jwt.Parse(tokenstr,func(t *jwt.Token) (any, error) {
		if _,ok:= t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil , errors.New("unexpected signature method")
		}
		return []byte(j.secretKey), nil
	},
	
) 
if err!= nil || !token.Valid{
  return nil,errors.New("invalid or expired token")
}

claims,ok:= token.Claims.(jwt.MapClaims)
if !ok {
	return nil,errors.New("error getting claims")
}


return &ports.TokenClaims{
	UserdID: int64(claims["sub"].(float64)), 
    Email: claims["email"].(string),
	Role: claims["role"].(string),} , nil
}

