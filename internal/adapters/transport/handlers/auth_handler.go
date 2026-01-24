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
	Auth(email string, password string)(string,error)
}

type FindProfile interface{
	FindMe(id int64)(*user.User,error)
}

type AuthHandler struct{
  userRegister UserRegisterer
  userLogin UserLogger
  userProfile FindProfile
}


func NewAuthHandler (user UserRegisterer, userLog UserLogger,UserMe FindProfile)*AuthHandler{
   return &AuthHandler{userRegister: user,userLogin:userLog ,}
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


func (au *AuthHandler) Login(c *gin.Context){
	var loginReq dto.LoginRequest

	if err:= c.ShouldBindJSON(&loginReq); err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid request"})
		return
	}

	u,err := au.userLogin.Auth(loginReq.Email,loginReq.Password); if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	
c.JSON(200, gin.H{
        "status":  "success",
        "message": "successful login",
        "token":   u, 
    })
}

func (au *AuthHandler)GetMe(c *gin.Context){
	idRaw,exist:= c.Get("id")
	if !exist{
		c.JSON(http.StatusForbidden,gin.H{"error":"id not found"})
		return
	}

	userId:= idRaw.(int64)

	u,err:=au.userProfile.FindMe(userId)
    if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"user not found"})
		return
	}

	response:= dto.UserResponse{
		Email: u.Email(),
		Role: u.Role(),
		Message: "WELCOME",
	}

	c.JSON(http.StatusOK,response)
	

}

	
 



