package repository

import (
	"Auth/internal/domain/user"
	"Auth/internal/ports"
	"errors"
	"sync"
)

type InMemoryStruct struct{
	users map[string]*user.User
	mu sync.RWMutex
}

var ErrorNotFound= errors.New("user not found")

func NewMemoryStruct()ports.UserRepository{
	return &InMemoryStruct{users: make(map[string]*user.User),}
}


func (in *InMemoryStruct) Save(user *user.User)error{
   in.mu.Lock()
   defer in.mu.Unlock()


   in.users[user.Email()]= user
   if user.GetId()==0{
      user.SetId(int64(len(in.users)+1))
   }
   return nil
}

func (in *InMemoryStruct) FindByEmail(email string)(*user.User,error){
   in.mu.RLock()
   defer in.mu.RUnlock()

   user, ok := in.users[email]
   if !ok {
	return nil, ErrorNotFound
   }
   return user ,nil
}