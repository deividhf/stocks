package main

import (
	"github.com/deividhf/stocks/controller"
	"github.com/deividhf/stocks/service"
	"github.com/gin-gonic/gin"
)

var (
	stockService    = service.New()
	stockController = controller.New(stockService)
)

func main() {
	server := gin.Default()

	server.GET("/stocks", func(ctx *gin.Context) {
		ctx.JSON(200, stockController.FindAll())
	})

	server.POST("/stocks", func(ctx *gin.Context) {
		ctx.JSON(201, stockController.Save(ctx))
	})

	server.Run(":8090")
}
