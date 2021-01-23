package repository

import (
	"log"

	"github.com/deividhf/stocks/stocks/entity"
	"gorm.io/gorm"
)

// StockRepository is responsible to manage the core logic by communicating with entity layer
type StockRepository interface {
	Save(stock entity.Stock) entity.Stock
	FindAll() []entity.Stock
}

type stockRepository struct {
	db *gorm.DB
}

func (s *stockRepository) Save(stock entity.Stock) entity.Stock {
	if result := s.db.Create(&stock); result.Error != nil {
		log.Panicf("Error on saving stock. %s", result.Error)
	}

	return stock
}

func (s *stockRepository) FindAll() []entity.Stock {
	stocks := make([]entity.Stock, 1)

	if result := s.db.Find(&stocks); result.Error != nil {
		log.Panicf("Error on getting all stocks. %s", result.Error)
	}

	return stocks
}

// New creates a new StockRepository
func New(db *gorm.DB) StockRepository {
	return &stockRepository{db}
}
