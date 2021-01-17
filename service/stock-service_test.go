package service

import (
	"testing"

	"github.com/deividhf/stocks/entity"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStockService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stock Service Test Suite")
}

var (
	weg = entity.Stock{"Weg", "WEGE3"}
)

var service StockService

var _ = Describe("Stock Service", func() {

	BeforeSuite(func() {
		service = New()
	})

	Describe("Fetching all stocks", func() {

		Context("When there is not stock", func() {

			BeforeEach(func() {
				service = New()
			})

			It("should return an empty array", func() {
				stocks := service.FindAll()
				Ω(stocks).Should(BeEmpty())
			})

		})

		Context("When it has stocks", func() {

			BeforeEach(func() {
				service.Save(weg)
			})

			It("array most not be empty", func() {
				stocks := service.FindAll()
				Ω(stocks).ShouldNot(BeEmpty())
			})

			It("should return the saved elements", func() {
				stock := service.FindAll()[0]
				Ω(stock).Should(Equal(weg))
			})

			AfterEach(func() {
				service = New()
			})
		})
	})

	Describe("Saving stocks", func() {

		Context("When add a stock", func() {

			BeforeEach(func() {
				service = New()
			})

			It("should return the saved element", func() {
				stock := service.Save(weg)
				Ω(stock).Should(Equal(weg))
			})
		})
	})
})
