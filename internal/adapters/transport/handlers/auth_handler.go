package transport

import (
	"Auth/internal/adapters/transport/dto"
	"Auth/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthHandler struct{
  userRegister *usecases.RegisterUserInput
}


func NewAuthHandler (user *usecases.RegisterUserInput)*AuthHandler{
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
	}

	
	c.JSON(http.StatusCreated,gin.H{"message":"user registered"})
}

	
 



