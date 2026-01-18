package transport

import (
	"Auth/internal/adapters/transport/dto"
	"Auth/internal/domain/user"
	"net/http"


	"github.com/gin-gonic/gin"
)

// create interface User register

type UserRegisterer interface{
	Execute(email string, password string)error
}

type UserLogger interface{
	Auth(email string, password string)(*user.User,error)
}

type AuthHandler struct{
  userRegister UserRegisterer
}


func NewAuthHandler (user UserRegisterer)*AuthHandler{
   return &AuthHandler{userRegister: user,}
}


func (au *AuthHandler) Register(c *gin.Context){
	var RegisterReq dto.RegisterRequest

	if err:= c.ShouldBindJSON(&RegisterReq);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid request"})
		return
	}

	if err := au.userRegister.Execute(RegisterReq.Email,RegisterReq.Password); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	
	c.JSON(http.StatusCreated,gin.H{"message":"user registered"})
}

	
 



