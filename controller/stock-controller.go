package controller

import (
	"github.com/deividhf/stocks/entity"
	"github.com/deividhf/stocks/service"
	"github.com/gin-gonic/gin"
)

// StockController deals with the requisition
type StockController interface {
	Save(ctx *gin.Context) error
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

func (c *stockController) Save(ctx *gin.Context) error {
	var stock entity.Stock

	if err := ctx.ShouldBindJSON(&stock); err != nil {
		return err
	}

	c.service.Save(stock)

	return nil
}

func (c *stockController) FindAll() []entity.Stock {
	return c.service.FindAll()
}
