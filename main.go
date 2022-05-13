package main

import (
	"loantracker-api/modules"

	"github.com/gin-gonic/gin"
)

func main () {
	modules.SayHi()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": modules.SayHi(),
		})
	})
	r.Run()
}
