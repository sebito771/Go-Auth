package main

import (
    "Auth/internal/adapters/repository"
    "Auth/internal/adapters/security"
    "Auth/internal/adapters/transport"
    handlers "Auth/internal/adapters/transport/handlers"
    "Auth/internal/usecases"
    
    "github.com/gin-gonic/gin"
)

func main(){
  //adapters
  repo:= repository.NewMemoryStruct()
  hasher:= security.BcryptStruct{}
  //use cases
  register:= usecases.NewRegisterUser(repo,&hasher)
  //handler
  handl:= handlers.NewAuthHandler(register)
   
  //gin init
  r:= gin.Default()

  //router
  transport.RegisterRoutes(r, handl)
  
  r.Run(":8000")
}