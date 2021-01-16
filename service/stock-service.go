package service

import "github.com/deividhf/stocks/entity"

type StockService interface {
	Save(stock entity.Stock) entity.Stock
	FindAll() []entity.Stock
}

type stockService struct {
	stocks []entity.Stock
}

func (s *stockService) Save(stock entity.Stock) entity.Stock {
	s.stocks = append(s.stocks, stock)
	return stock
}

func (s *stockService) FindAll() []entity.Stock {
	return s.stocks
}

// New create a new StockService
func New() StockService {
	return &stockService{stocks: []entity.Stock{}}
}
