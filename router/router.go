package router

import (
	"loantracker-api/router/routes"

	"github.com/gin-gonic/gin"
)

func Run() {

	router := gin.Default()
	
	router.GET("/", routes.Hello)

	router.Run()
}
