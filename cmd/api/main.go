package main

import (
	
	 "Auth/internal/adapters/transport"
	"github.com/gin-gonic/gin"
)


func main(){
  r:= gin.Default() 
  transport.RegisterRoutes(r)
  r.Run(":8000")
}