package usecases

import (
	"Auth/internal/ports"
	"errors"
	"strings"
	"time"
)

type LogoutUser struct {
	blackList ports.TokenBlackList
	tokenValidator ports.TokenValidator
}

func NewLogoutUser(blackList ports.TokenBlackList, tokenValidator ports.TokenValidator) *LogoutUser {
	return &LogoutUser{blackList: blackList, tokenValidator: tokenValidator}
}


func (l *LogoutUser) Logout(token string) error{
	if strings.TrimSpace(token) == "" {
    return errors.New("invalid token: token cannot be empty")
   }

  claims,err:= l.tokenValidator.ValidateToken(token)
  if err != nil {
	return errors.New("invalid token: " + err.Error())
  }

  ttl:= int64(time.Until(time.Unix(claims.Exp,0)).Seconds())
  if ttl <= 0 {
	return errors.New("token already expired")
  }

   if exist,err:= l.blackList.IsBlackListed(token); err != nil {
	return errors.New("error checking token black list: " + err.Error())
   }else if exist{
	return errors.New("token already blacklisted")
   }
	return l.blackList.Add(token, ttl)
}