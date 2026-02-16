package transport

import (
	"Auth/internal/adapters/transport/dto"
	"Auth/internal/domain/user"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

// create interface User register

type UseCaseRegister interface{
	Execute(email string, password string)error
}

type UseCaseLogin interface{
	Auth(email string, password string)(string,error)
}

type ProfileUseCase interface{
	FindMe(id int64)(*user.User,error)
}

type LogoutUseCase interface{
	Logout(token string)error
}

type AuthHandler struct{
  userRegister UseCaseRegister
  userLogin UseCaseLogin
  userProfile ProfileUseCase
  userLogout LogoutUseCase
}


func NewAuthHandler (user UseCaseRegister, userLog UseCaseLogin,UserMe ProfileUseCase,userLogout LogoutUseCase)*AuthHandler{
   return &AuthHandler{userRegister: user,userLogin:userLog ,userProfile: UserMe,userLogout:userLogout}
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

func (au *AuthHandler) Logout(c *gin.Context){
  authHeader:= c.GetHeader("Authorization")

  if authHeader==""{
	c.JSON(http.StatusUnauthorized,gin.H{"error":"token required"})
	return 
  }
  if !strings.HasPrefix(authHeader,"Bearer "){
	c.JSON(http.StatusUnauthorized,gin.H{"error":"invalid token"})
	return 
  } 

  TokenStr:= strings.TrimPrefix(authHeader,"Bearer ")
  if err:= au.userLogout.Logout(TokenStr); err != nil{
	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	return 
  }
    c.JSON(http.StatusAccepted,gin.H{"message":"successful logout"})
}
	
 



