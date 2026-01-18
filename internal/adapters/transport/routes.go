package transport

import (
	"github.com/gin-gonic/gin"
	handlers "Auth/internal/adapters/transport/handlers"
	
)


func RegisterRoutes(r *gin.Engine, h *handlers.AuthHandler) {

	// create auth group
    
	auth:= r.Group("/auth")
	{
       auth.POST("/register",h.Register)
	   auth.POST("/login",h.Login)
	}
	// health check route
 
	r.GET("/health",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message":"server running",
		})
	})
}