package repository

import (
	"Auth/internal/domain/user"
	"Auth/internal/ports"
	"errors"
	"strconv"
	"sync"
)

type InMemoryStruct struct{
	users map[string]*user.User
   usersByid map[string]*user.User
	mu sync.RWMutex
}

var ErrorNotFound= errors.New("user not found")

func NewMemoryStruct()ports.UserRepository{
	return &InMemoryStruct{users: make(map[string]*user.User),usersByid: make(map[string]*user.User)}
}


func (in *InMemoryStruct) Save(user *user.User)error{
   in.mu.Lock()
   defer in.mu.Unlock()

   if user.GetId()==0{
      user.SetId(int64(len(in.users)+1))
   }
   in.users[user.Email()]= user
   in.usersByid[strconv.FormatInt(user.GetId(),10)]= user
  
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

func (in *InMemoryStruct) FindById(id int64)(*user.User,error){
   in.mu.RLock()
   defer in.mu.RUnlock()
   user,ok:= in.usersByid[strconv.FormatInt(id, 10)]
   if !ok{
      return nil,ErrorNotFound
   }
   return user,nil
}