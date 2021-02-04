package repository

import (
	"log"

	"github.com/deividhf/stocks/stocks/entity"
	"gorm.io/gorm"
)

// StockRepository is responsible to manage the core logic by communicating with entity layer
type StockDatabaseRepository struct {
	db *gorm.DB
}

func (r *StockDatabaseRepository) Save(stock entity.Stock) entity.Stock {
	if result := r.db.Create(&stock); result.Error != nil {
		log.Panicf("Error on saving stock. %s", result.Error)
	}

	return stock
}

func (r *StockDatabaseRepository) FindAll() []entity.Stock {
	stocks := make([]entity.Stock, 1)

	if result := r.db.Find(&stocks); result.Error != nil {
		log.Panicf("Error on getting all stocks. %s", result.Error)
	}

	return stocks
}

func (r *StockDatabaseRepository) GetByID(id string) (entity.Stock, error) {
	stock := entity.Stock{}

	if err := r.db.First(&stock, id).Error; err != nil {
		return stock, err
	}

	return stock, nil
}

func (r *StockDatabaseRepository) DeleteByID(id string) error {
	if r.db.Delete(&entity.Stock{}, id).RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// New creates a new StockRepository
func New(db *gorm.DB) *StockDatabaseRepository {
	return &StockDatabaseRepository{db}
}
