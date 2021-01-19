package mocks

import (
	"github.com/deividhf/stocks/stocks/entity"
	"github.com/stretchr/testify/mock"
)

// StockServiceMock is used on test
type StockServiceMock struct {
	mock.Mock
}

// Save is a mock method
func (m *StockServiceMock) Save(stock entity.Stock) entity.Stock {
	args := m.Called(stock)
	return args.Get(0).(entity.Stock)
}

// FindAll is a mock method
func (m *StockServiceMock) FindAll() []entity.Stock {
	args := m.Called()
	return args.Get(0).([]entity.Stock)
}
