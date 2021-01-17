package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var router *gin.Engine
var w *httptest.ResponseRecorder

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Test Suite")
}

var _ = Describe("Server Test", func() {

	BeforeSuite(func() {
		router = setupServer()
	})

	Describe("Getting stocks", func() {

		Context("When there is no stocks", func() {

			BeforeEach(func() {
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

			var savedStock string

			BeforeEach(func() {
				// It will be replaced with mocks
				w = httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/stocks", strings.NewReader(`{"name":"Weg","ticker":"WEGE3"}`))
				router.ServeHTTP(w, req)
				savedStock = w.Body.String()

				w = httptest.NewRecorder()
				req, _ = http.NewRequest("GET", "/stocks", nil)
				router.ServeHTTP(w, req)
			})

			It("should return a saved stock", func() {
				expected := fmt.Sprintf("[%s]", savedStock)
				Ω(w.Body.String()).Should(Equal(expected))
			})
		})

	})

	Describe("Posting stocks", func() {

		Context("Saving a new stock", func() {

			BeforeEach(func() {
				w = httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/stocks", strings.NewReader(`{"name":"Weg","ticker":"WEGE3"}`))
				router.ServeHTTP(w, req)
			})

			It("should return the saved stock", func() {
				expected := `{"name":"Weg","ticker":"WEGE3"}`
				Ω(w.Body.String()).Should(Equal(expected))
				Ω(w.Code).Should(Equal(201))
			})
		})
	})
})
