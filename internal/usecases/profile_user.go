package usecases

import (
	"Auth/internal/domain/user"
	"Auth/internal/ports"
)

type ProfileUser struct {
	 repo ports.UserRepository
}

func NewProfilUser(repo ports.UserRepository)*ProfileUser{
	return &ProfileUser{repo: repo}
}

func (p *ProfileUser)FindMe(id int64)(*user.User,error){
	u,err:= p.repo.FindById(id)
	if err!=nil{
		return nil,err
	}
	return u,nil
}