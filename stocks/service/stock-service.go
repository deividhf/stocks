package service

import (
	"log"

	"github.com/deividhf/stocks/stocks/entity"
	"gorm.io/gorm"
)

// StockService is responsible to manage the core logic by communicating with entity layer
type StockService interface {
	Save(stock entity.Stock) entity.Stock
	FindAll() []entity.Stock
}

type stockService struct {
	db *gorm.DB
}

func (s *stockService) Save(stock entity.Stock) entity.Stock {
	if result := s.db.Create(&stock); result.Error != nil {
		log.Panicf("Error on saving stock. %s", result.Error)
	}

	return stock
}

func (s *stockService) FindAll() []entity.Stock {
	stocks := make([]entity.Stock, 1)

	if result := s.db.Find(&stocks); result.Error != nil {
		log.Panicf("Error on getting all stocks. %s", result.Error)
	}

	return stocks
}

// New creates a new StockService
func New(db *gorm.DB) StockService {
	return &stockService{db}
}
