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
	GetByID(id string) (entity.Stock, error)
	DeleteByID(id string) error
}

type stockRepository struct {
	db *gorm.DB
}

func (r *stockRepository) Save(stock entity.Stock) entity.Stock {
	if result := r.db.Create(&stock); result.Error != nil {
		log.Panicf("Error on saving stock. %s", result.Error)
	}

	return stock
}

func (r *stockRepository) FindAll() []entity.Stock {
	stocks := make([]entity.Stock, 1)

	if result := r.db.Find(&stocks); result.Error != nil {
		log.Panicf("Error on getting all stocks. %s", result.Error)
	}

	return stocks
}

func (r *stockRepository) GetByID(id string) (entity.Stock, error) {
	stock := entity.Stock{}

	if err := r.db.First(&stock, id).Error; err != nil {
		return stock, err
	}

	return stock, nil
}

func (r *stockRepository) DeleteByID(id string) error {
	if r.db.Delete(&entity.Stock{}, id).RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// New creates a new StockRepository
func New(db *gorm.DB) StockRepository {
	return &stockRepository{db}
}
