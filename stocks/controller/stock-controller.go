package controller

import (
	"github.com/deividhf/stocks/stocks/entity"
	"github.com/deividhf/stocks/stocks/repository"
	"github.com/gin-gonic/gin"
)

// StockController deals with the request
type StockController interface {
	Save(ctx *gin.Context) (entity.Stock, error)
	FindAll() []entity.Stock
}

type stockController struct {
	repository repository.StockRepository
}

// New creates a StockController
func New(repository repository.StockRepository) StockController {
	return &stockController{
		repository: repository,
	}
}

func (c *stockController) Save(ctx *gin.Context) (entity.Stock, error) {
	var stock entity.Stock

	if err := ctx.ShouldBindJSON(&stock); err != nil {
		return stock, err
	}

	return c.repository.Save(stock), nil
}

func (c *stockController) FindAll() []entity.Stock {
	return c.repository.FindAll()
}
