package controller

import (
	"github.com/deividhf/stocks/entity"
	"github.com/deividhf/stocks/service"
	"github.com/gin-gonic/gin"
)

type StockController interface {
	Save(ctx *gin.Context) entity.Stock
	FindAll() []entity.Stock
}

type stockController struct {
	service service.StockService
}

// New creates a StockController
func New(service service.StockService) StockController {
	return &stockController{
		service: service,
	}
}

func (c *stockController) Save(ctx *gin.Context) entity.Stock {
	var stock entity.Stock
	ctx.BindJSON(&stock)

	return c.service.Save(stock)
}

func (c *stockController) FindAll() []entity.Stock {
	return c.service.FindAll()
}
