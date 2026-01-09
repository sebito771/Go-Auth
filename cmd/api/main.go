package main

import (
	"github.com/gin-gonic/gin"
  "Auth/internal/adapter/HTTP"

)


func main(){
  r:= gin.Default() 
  http.RegisterRoutes(r)
  r.Run(":8000")
}