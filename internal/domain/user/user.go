package user

import (
	"errors"
	"strings"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	id int64
	email string
	role  Role
	password string
}


func New(email string, password string) (*User, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	if email == "" {
		return nil, errors.New("email is required")
	}

	if password == ""{
		return nil , errors.New("password is required")
	}

	return &User{
		email: email,
		role:  RoleUser,
		password: password,
	}, nil
}


func (u *User) PromoteToAdmin() {
	u.role = RoleAdmin
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Role() Role {
	return u.role
}


func (u *User) Password() string{
	return  u.password
}

func (u *User) GetId() int64{
	return u.id
}

func (u *User) SetId(id int64){
  u.id=id
}