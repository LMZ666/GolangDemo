package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	var1 := "abcd"
	fmt.Println(var1)
	r.Run() // LISTEN AND SERVE ON 0.0.0.0:8080
}
