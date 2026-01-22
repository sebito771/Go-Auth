package usecases

import (
	"Auth/internal/ports"
	"errors"
)

type LoginUser struct{
	repo ports.UserRepository
	hasher ports.PassWordHaser
	tokenGen ports.TokenGenerator
}

var invalidCredentials = errors.New("invalid credentials")

func NewLoginUser(repo ports.UserRepository , hasher ports.PassWordHaser,token ports.TokenGenerator) *LoginUser{
	return &LoginUser{repo: repo, hasher: hasher,tokenGen: token}
}



func (l *LoginUser)Auth(email string, password string)(string,error){

	existingEmail,err:= l.repo.FindByEmail(email)
    
	 if err != nil{
		return "" , invalidCredentials
	 }

	 hash:= existingEmail.Password()
	 err= l.hasher.Compare(password,hash)
	 if err != nil{
		return "" , invalidCredentials
	 }

	 token,err:= l.tokenGen.GetToken(existingEmail)
	 if err != nil{
		return "" , err
	 }

	 return token , nil
}