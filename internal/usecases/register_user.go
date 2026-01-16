package usecases


import (
	"errors"
	"Auth/internal/ports"
	"Auth/internal/domain/user"
	"Auth/internal/adapters/repository"
)

var EmailAlreadyExist = errors.New("email already exist")


type RegisterUserInput struct{
	repo ports.UserRepository
	hasher ports.PassWordHaser
}


func NewRegisterUser(repo ports.UserRepository, hasher ports.PassWordHaser) *RegisterUserInput{
	return &RegisterUserInput{repo: repo,hasher: hasher}
	
}


func (uc RegisterUserInput) Execute(email string, password string ) error{
     
	// verify if the email already exist
	 existingEmail,err := uc.repo.FindByEmail(email)
	 if err!= nil&& !errors.Is(err,repository.ErrorNotFound){
		return err
	 }
	 if existingEmail != nil{
		return EmailAlreadyExist
	 }

	// hash password

	passwordHash,err := uc.hasher.Hash(password)
	if err != nil{
		return err
	}
   
    // create new User
	newUser,err:= user.New(email,passwordHash)
    if err != nil{
		return err
	}

	//save User domain
	if err := uc.repo.Save(newUser); err!=nil{
		return err
	}
	 
	 return nil 
}