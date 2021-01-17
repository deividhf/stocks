package main

import (
	"net/http"

	"github.com/deividhf/stocks/controller"
	"github.com/deividhf/stocks/service"
	"github.com/gin-gonic/gin"
)

var (
	stockService    = service.New()
	stockController = controller.New(stockService)
)

func main() {
	setupServer().Run(":8090")
}

func setupServer() *gin.Engine {
	server := gin.Default()

	server.GET("/stocks", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, stockController.FindAll())
	})

	server.POST("/stocks", func(ctx *gin.Context) {
		stock, err := stockController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusCreated, stock)
		}
	})

	return server
}
