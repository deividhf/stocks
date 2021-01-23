package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/deividhf/stocks/stocks/entity"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestStockRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stock Repository Test Suite")
}

var (
	weg        = entity.Stock{Name: "Weg", Ticker: "WEGE3"}
	repository StockRepository
	db         *sql.DB
	gdb        *gorm.DB
	mock       sqlmock.Sqlmock
)

var _ = Describe("Stock Repository", func() {

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		Ω(err).ShouldNot(HaveOccurred())

		dialector := sqlite.Dialector{
			DriverName: sqlite.DriverName,
			DSN:        "sqlmock_db_0",
			Conn:       db,
		}

		gdb, err = gorm.Open(dialector, &gorm.Config{})
		Ω(err).ShouldNot(HaveOccurred())

		repository = New(gdb)
	})

	Describe("Fetching all stocks", func() {

		Context("When there is not stock", func() {

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"name", "ticker"})
				mock.ExpectQuery("SELECT \\* FROM `stocks`").WillReturnRows(rows)
			})

			It("should return an empty array", func() {
				stocks := repository.FindAll()
				Ω(stocks).Should(BeEmpty())
			})

		})

		Context("When it has stocks", func() {

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"name", "ticker"}).AddRow("Weg", "WEGE3")
				mock.ExpectQuery("SELECT \\* FROM `stocks`").WillReturnRows(rows)
			})

			It("array most not be empty", func() {
				stocks := repository.FindAll()
				Ω(stocks).ShouldNot(BeEmpty())
			})

			It("should return the saved elements", func() {
				stock := repository.FindAll()[0]
				Ω(stock).Should(Equal(weg))
			})

		})
	})

	Describe("Saving stocks", func() {

		Context("When add a stock", func() {

			BeforeEach(func() {
				mock.ExpectExec("INSERT INTO `stocks`").WithArgs("Weg", "WEGE3").WillReturnResult(sqlmock.NewResult(0, 1))
			})

			It("should return the saved element", func() {
				stock := repository.Save(weg)
				Ω(stock).Should(Equal(weg))
			})
		})
	})

	AfterEach(func() {
		db.Close()
	})
})