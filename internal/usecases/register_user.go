package usecases


import (
	"errors"
	"Auth/internal/ports"
	"Auth/internal/domain/user"
)

var EmailAlreadyExist = errors.New("email already exist")


type RegisterUserInput struct{
	repo ports.UserRepository
}


func NewRegisterUser(repo ports.UserRepository) *RegisterUserInput{
	return &RegisterUserInput{repo: repo,}
}


func (uc RegisterUserInput) excute(email string ) error{
     
	// verify if the email already exist
	 existingEmail,_ := uc.repo.FindByEmail(email)
	 if existingEmail != nil{
		return EmailAlreadyExist
	 }


    // create new User
	newUser,err:= user.New(email)
    if err != nil{
		return err
	}

	//save User domain
	if err := uc.repo.Save(newUser); err!=nil{
		return err
	}
	 
	 return nil 
}