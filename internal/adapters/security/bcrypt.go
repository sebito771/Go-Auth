package security

import "golang.org/x/crypto/bcrypt"


type BcryptStruct struct{}


func (b *BcryptStruct) Hash(password string)(string,error){
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes),err
}