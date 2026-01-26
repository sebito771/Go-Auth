package transport

import (
	handlers "Auth/internal/adapters/transport/handlers"
	"Auth/internal/adapters/transport/middlewares"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(r *gin.Engine, h *handlers.AuthHandler,m middlewares.AuthMiddleWare) {

	// create auth group
    
	auth:= r.Group("/auth")
	{
       auth.POST("/register",h.Register)
	   auth.POST("/login",h.Login)
	}

	user:= r.Group("/logged")
	{
		user.Use(m.Aunthenticate())

		user.GET("/me",h.GetMe)
	}
	// health check route
 
	r.GET("/health",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message":"server running",
		})
	})
}