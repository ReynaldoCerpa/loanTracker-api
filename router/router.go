package router

import (
	"loantracker-api/router/routes"

	"github.com/gin-gonic/gin"
)

func Run() {

	router := gin.Default()
	
	router.GET("/hello", routes.Hello)

	router.Run()
}
