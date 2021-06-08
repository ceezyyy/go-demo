package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// *Engine
	r := gin.Default()

	r.GET("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.Run() // listen and serve on localhost:8080

}
