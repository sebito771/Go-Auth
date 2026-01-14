package transport

import (
	"github.com/gin-gonic/gin"
	
)


func RegisterRoutes(r *gin.Engine) {

  
	r.GET("/health",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message":"server running",
		})
	})
}