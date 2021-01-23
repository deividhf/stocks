package stocks

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/deividhf/stocks/stocks/controller"
	"github.com/deividhf/stocks/stocks/entity"
	"github.com/deividhf/stocks/stocks/service/mocks"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stocks Router Test Suite")
}

var router *gin.Engine
var w *httptest.ResponseRecorder
var serviceMock *mocks.StockServiceMock

var _ = Describe("Stocks Router Test", func() {

	BeforeSuite(func() {
		router = gin.Default()
		serviceMock = &mocks.StockServiceMock{}

		stockRouter := New(controller.New(serviceMock))
		stockRouter.Routes(router)
	})

	Describe("Getting stocks", func() {

		Context("When there is no stocks", func() {

			BeforeEach(func() {
				serviceMock.On("FindAll").Return([]entity.Stock{}).Once()

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

			var stock entity.Stock

			BeforeEach(func() {
				stock = entity.Stock{ID: 0, Name: "Weg", Ticker: "WEGE3"}
				serviceMock.On("FindAll").Return([]entity.Stock{stock}).Once()

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

	Describe("Posting stocks", func() {

		Context("Saving a new stock", func() {

			BeforeEach(func() {
				savedStock := entity.Stock{ID: 0, Name: "Weg", Ticker: "WEGE3"}
				serviceMock.On("Save", savedStock).Return(savedStock).Once()

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
})
