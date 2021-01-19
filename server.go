package main

import (
	"github.com/deividhf/stocks/stocks"
	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run(":8090")
}

func setupServer() *gin.Engine {
	server := gin.Default()

	router := stocks.DefaultRouter()
	router.Routes(server)

	return server
}
