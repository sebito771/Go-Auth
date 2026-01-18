package usecases

import (
	"Auth/internal/ports"
	"errors"
	"Auth/internal/domain/user"
)

type LoginUser struct{
	repo ports.UserRepository
	hasher ports.PassWordHaser
}

var invalidCredentials = errors.New("invalid credentials")

func NewLoginUser(repo ports.UserRepository , hasher ports.PassWordHaser) *LoginUser{
	return &LoginUser{repo: repo, hasher: hasher}
}



func (l *LoginUser)Auth(email string, password string)(*user.User,error){

	existingEmail,err:= l.repo.FindByEmail(email)
    
	 if err != nil{
		return nil , invalidCredentials
	 }

	 hash:= existingEmail.Password()
	 err= l.hasher.Compare(password,hash)
	 if err != nil{
		return nil , invalidCredentials
	 }
	 return existingEmail , nil
}