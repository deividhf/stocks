package controller

import (
	"github.com/deividhf/stocks/stocks/entity"
	"github.com/deividhf/stocks/stocks/service"
	"github.com/gin-gonic/gin"
)

// StockController deals with the request
type StockController interface {
	Save(ctx *gin.Context) (entity.Stock, error)
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

func (c *stockController) Save(ctx *gin.Context) (entity.Stock, error) {
	var stock entity.Stock

	if err := ctx.ShouldBindJSON(&stock); err != nil {
		return stock, err
	}

	return c.service.Save(stock), nil
}

func (c *stockController) FindAll() []entity.Stock {
	return c.service.FindAll()
}
