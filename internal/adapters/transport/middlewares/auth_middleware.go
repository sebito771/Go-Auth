package middlewares

import (
	"Auth/internal/ports"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleWare struct{
	validator ports.TokenValidator
}

func NewAuhtMiddleWare(vdt ports.TokenValidator)*AuthMiddleWare{
	return &AuthMiddleWare{validator: vdt}
}


func (am *AuthMiddleWare) Aunthenticate()gin.HandlerFunc {
 return func(ctx *gin.Context) {
	authHeader:= ctx.GetHeader("Authorization")

	if authHeader==""{
		ctx.JSON(http.StatusUnauthorized,gin.H{"error":"token required"})
		ctx.Abort()
		return 
	}

	if !strings.HasPrefix(authHeader,"Bearer "){
		ctx.JSON(http.StatusUnauthorized,gin.H{"error":"invalid token"})
		return 
	}

	TokenStr:= strings.TrimPrefix(authHeader,"Bearer ")

	tk,err:=am.validator.ValidateToken(TokenStr)
	if err!=nil{
	 ctx.JSON(http.StatusUnauthorized,gin.H{"error":"invalid or expired token"})
	 ctx.Abort()
	 return 
	}
	
	ctx.Set("id",tk.UserdID)
	ctx.Set("role",tk.Role)

	ctx.Next()

 }
}
