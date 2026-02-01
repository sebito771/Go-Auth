package main

import (
	"Auth/internal/adapters/repository/mariadb"
	"Auth/internal/adapters/security"
	"Auth/internal/adapters/transport"
	handlers "Auth/internal/adapters/transport/handlers"
	"Auth/internal/adapters/transport/middlewares"
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
    host:= os.Getenv("DB_HOST")
    port:= os.Getenv("DB_PORT")
    user:= os.Getenv("DB_USER")
    password:= os.Getenv("DB_PASSWORD")
    dbname:= os.Getenv("DB_NAME")

  //adapters
  repoMaria, err:= mariadb.NewMariaDBRepo(user,password,host+":"+port,dbname)
  if err != nil{
  log.Fatalf("Error connecting to MariaDB: %v", err) 
  }
  
  //repo:= repository.NewMemoryStruct()
  hasher:= security.BcryptStruct{}
  tokenGen:= security.NewJwtAdapter(tokenKey)
  
  //use cases
  register:= usecases.NewRegisterUser(repoMaria,&hasher)
  login:= usecases.NewLoginUser(repoMaria,&hasher,tokenGen)
  profile:= usecases.NewProfilUser(repoMaria)
  //handler
  handl:= handlers.NewAuthHandler(register,login,profile)
  //middlewares
  middl:= middlewares.NewAuhtMiddleWare(tokenGen)
   
  //gin init
  r:= gin.Default()

  //router
  transport.RegisterRoutes(r, handl,*middl)
  
  r.Run(":8000")
}