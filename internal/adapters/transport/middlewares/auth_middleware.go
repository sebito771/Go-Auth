package middlewares

import (
	"Auth/internal/ports"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleWare struct{
	validator ports.TokenValidator
	blackList ports.TokenBlackList
}

func NewAuthMiddleWare(vdt ports.TokenValidator,bl ports.TokenBlackList)*AuthMiddleWare{
	return &AuthMiddleWare{validator: vdt,blackList:bl}
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
		ctx.JSON(http.StatusUnauthorized,gin.H{"error":"invalid token format"})
		ctx.Abort()
		return 
	}

	TokenStr:= strings.TrimPrefix(authHeader,"Bearer ")

	tk,err:=am.validator.ValidateToken(TokenStr)
	if err!=nil{
	 ctx.JSON(http.StatusUnauthorized,gin.H{"error":"invalid or expired token"})
	 ctx.Abort()
	 return 
	}

    exist,err := am.blackList.IsBlackListed(TokenStr)
	
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":"error checking token blacklist"})
		ctx.Abort()
		return
	}
	
	if exist{
		ctx.JSON(http.StatusUnauthorized,gin.H{"error":"token blacklisted"})
		ctx.Abort()
		return
	}

	
	ctx.Set("id",tk.UserdID)
	ctx.Set("role",tk.Role)

	ctx.Next()

 }
}
