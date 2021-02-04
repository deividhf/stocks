package stocks

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/deividhf/stocks/stocks/controller"
	"github.com/deividhf/stocks/stocks/entity"
	"github.com/deividhf/stocks/stocks/repository/mocks"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stocks Router Test Suite")
}

var router *gin.Engine
var w *httptest.ResponseRecorder
var repositoryMock *mocks.StockRepository

var _ = BeforeSuite(func() {
	router = gin.Default()
	repositoryMock = &mocks.StockRepository{}

	stockRouter := New(controller.New(repositoryMock))
	stockRouter.Routes(router)
})

var _ = Describe("Stocks Router Test", func() {

	var stock = entity.Stock{ID: 0, Name: "Weg", Ticker: "WEGE3"}

	Describe("Getting stocks", func() {

		Context("When there is no stocks", func() {

			BeforeEach(func() {
				repositoryMock.On("FindAll").Return([]entity.Stock{}).Once()

				w = httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/stocks", nil)
				router.ServeHTTP(w, req)
			})

			It("should return an empty array", func() {
				expected := "[]"
				Ω(w.Body.String()).Should(Equal(expected))
				Ω(w.Code).Should(Equal(200))
			})

		})

		Context("When there are stocks", func() {

			BeforeEach(func() {
				repositoryMock.On("FindAll").Return([]entity.Stock{stock}).Once()

				w = httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/stocks", nil)
				router.ServeHTTP(w, req)
			})

			It("should return a saved stock", func() {
				expected := `[{"id":0,"name":"Weg","ticker":"WEGE3"}]`
				Ω(w.Body.String()).Should(Equal(expected))
				Ω(w.Code).Should(Equal(200))

			})
		})

	})

	Describe("Getting stock by id", func() {

		Context("When there is no stock", func() {

			BeforeEach(func() {
				repositoryMock.On("GetByID", "123").Return(stock, gorm.ErrRecordNotFound).Once()

				w = httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/stocks/123", nil)
				router.ServeHTTP(w, req)
			})

			It("should return not found", func() {
				Ω(w.Code).Should(Equal(404))
				Ω(w.Body.String()).Should(Equal(`{"error":"record not found"}`))
			})
		})

		Context("When there is a stock", func() {

			BeforeEach(func() {
				repositoryMock.On("GetByID", "0").Return(stock, nil).Once()

				w = httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/stocks/0", nil)
				router.ServeHTTP(w, req)
			})

			It("should return the stock", func() {
				expected := `{"id":0,"name":"Weg","ticker":"WEGE3"}`
				Ω(w.Code).Should(Equal(200))
				Ω(w.Body.String()).Should(Equal(expected))
			})
		})
	})

	Describe("Posting stocks", func() {

		Context("Saving a new stock", func() {

			BeforeEach(func() {
				savedStock := entity.Stock{ID: 0, Name: "Weg", Ticker: "WEGE3"}
				repositoryMock.On("Save", savedStock).Return(savedStock).Once()

				w = httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/stocks", strings.NewReader(`{"name":"Weg","ticker":"WEGE3"}`))
				router.ServeHTTP(w, req)
			})

			It("should return the saved stock", func() {
				expected := `{"id":0,"name":"Weg","ticker":"WEGE3"}`
				Ω(w.Body.String()).Should(Equal(expected))
				Ω(w.Code).Should(Equal(201))
			})
		})
	})

	Describe("Deleting stocks", func() {

		Context("When there is no stock", func() {

			BeforeEach(func() {
				repositoryMock.On("DeleteByID", "123").Return(gorm.ErrRecordNotFound).Once()

				w = httptest.NewRecorder()
				req, _ := http.NewRequest("DELETE", "/stocks/123", nil)
				router.ServeHTTP(w, req)
			})

			It("should return not found", func() {
				Ω(w.Code).Should(Equal(404))
				Ω(w.Body.String()).Should(Equal(`{"error":"record not found"}`))
			})
		})

		Context("When there is a stock", func() {

			BeforeEach(func() {
				repositoryMock.On("DeleteByID", "123").Return(nil).Once()

				w = httptest.NewRecorder()
				req, _ := http.NewRequest("DELETE", "/stocks/123", nil)
				router.ServeHTTP(w, req)
			})

			It("should delete the record", func() {
				Ω(w.Code).Should(Equal(204))
				Ω(w.Body.String()).Should(BeEmpty())
			})
		})
	})
})
