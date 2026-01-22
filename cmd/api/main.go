package main

import (
	"Auth/internal/adapters/repository"
	"Auth/internal/adapters/security"
	"Auth/internal/adapters/transport"
	handlers "Auth/internal/adapters/transport/handlers"
	"Auth/internal/usecases"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
    err:= godotenv.Load(); if err != nil{
	 log.Fatalf("Error loading .env file: %v", err) 
	 }

     tokenKey:= os.Getenv("TOKEN_PASSWORD")
  //adapters
  repo:= repository.NewMemoryStruct()
  hasher:= security.BcryptStruct{}
  tokenGen:= security.NewJwtAdapter(tokenKey)
  //use cases
  register:= usecases.NewRegisterUser(repo,&hasher)
  login:= usecases.NewLoginUser(repo,&hasher,tokenGen)
  //handler
  handl:= handlers.NewAuthHandler(register,login)
   
  //gin init
  r:= gin.Default()

  //router
  transport.RegisterRoutes(r, handl)
  
  r.Run(":8000")
}