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

type Config struct {
	TokenKey  string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPassword string
	DBName    string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := Config{
		TokenKey:  os.Getenv("TOKEN_PASSWORD"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:    os.Getenv("DB_NAME"),
	}

	//adapters
	repoMaria, err := mariadb.NewMariaDBRepo(config.DBUser, config.DBPassword, config.DBHost+":"+config.DBPort, config.DBName)
	if err != nil {
		log.Fatalf("Error connecting to MariaDB: %v", err)
	}

	blacklist := security.NewBlackList()
	hasher := security.BcryptStruct{}
	tokenGen := security.NewJwtAdapter(config.TokenKey)

	//use cases
	register := usecases.NewRegisterUser(repoMaria, &hasher)
	login := usecases.NewLoginUser(repoMaria, &hasher, tokenGen)
	profile := usecases.NewProfilUser(repoMaria)
	logout := usecases.NewLogoutUser(blacklist, tokenGen)
	refresh := usecases.NewRefreshTokenUseCase(repoMaria, tokenGen, tokenGen, blacklist)
	//handler
	handl := handlers.NewAuthHandler(register, login, profile, logout, refresh)
	//middlewares
	middl := middlewares.NewAuthMiddleWare(tokenGen, blacklist)

	//gin init
	r := gin.Default()

	//router
	transport.RegisterRoutes(r, handl, *middl)

	r.Run(":8000")
}
