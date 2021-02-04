package controller

import (
	"github.com/deividhf/stocks/stocks/entity"
	"github.com/gin-gonic/gin"
)

type StockRepository interface {
	Save(stock entity.Stock) entity.Stock
	FindAll() []entity.Stock
	GetByID(id string) (entity.Stock, error)
	DeleteByID(id string) error
}

type StockResourceController struct {
	repository StockRepository
}

// New creates a StockController
func New(repository StockRepository) *StockResourceController {
	return &StockResourceController{
		repository: repository,
	}
}

func (c *StockResourceController) Save(ctx *gin.Context) (entity.Stock, error) {
	var stock entity.Stock

	if err := ctx.ShouldBindJSON(&stock); err != nil {
		return stock, err
	}

	return c.repository.Save(stock), nil
}

func (c *StockResourceController) FindAll() []entity.Stock {
	return c.repository.FindAll()
}

func (c *StockResourceController) GetByID(id string) (entity.Stock, error) {
	return c.repository.GetByID(id)
}
