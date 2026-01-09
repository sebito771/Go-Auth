package main

import (
	"github.com/gin-gonic/gin"
)


func main(){
  r:= gin.Default() 

  r.GET("/ping",func(ctx *gin.Context) {
	  ctx.JSON(200,gin.H{
		"mensaje":"servidor vivo",
	  })

  })

  r.Run(":8000")
}