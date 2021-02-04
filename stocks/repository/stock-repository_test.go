package repository

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/deividhf/stocks/stocks/entity"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestStockRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stock Repository Test Suite")
}

var (
	weg        = entity.Stock{ID: 123, Name: "Weg", Ticker: "WEGE3"}
	repository *StockDatabaseRepository
	db         *sql.DB
	gdb        *gorm.DB
	mock       sqlmock.Sqlmock
)

var _ = Describe("Stock Repository", func() {

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		Ω(err).ShouldNot(HaveOccurred())

		rows := sqlmock.NewRows([]string{"VERSION()"}).AddRow("3.7")
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		dialector := mysql.New(mysql.Config{
			Conn: db,
		})

		gdb, err = gorm.Open(dialector, &gorm.Config{})
		Ω(err).ShouldNot(HaveOccurred())

		repository = New(gdb)
	})

	Describe("Fetching all stocks", func() {

		Context("When there is not stock", func() {

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "name", "ticker"})
				mock.ExpectQuery("SELECT \\* FROM `stocks`").WillReturnRows(rows)
			})

			It("should return an empty array", func() {
				stocks := repository.FindAll()
				Ω(stocks).Should(BeEmpty())
				Ω(mock.ExpectationsWereMet()).Should(BeNil())
			})

		})

		Context("When it has stocks", func() {

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "name", "ticker"}).AddRow("123", "Weg", "WEGE3")
				mock.ExpectQuery("SELECT \\* FROM `stocks`").WillReturnRows(rows)
			})

			It("array most not be empty", func() {
				stocks := repository.FindAll()
				Ω(stocks).ShouldNot(BeEmpty())
				Ω(mock.ExpectationsWereMet()).Should(BeNil())
			})

			It("should return the saved elements", func() {
				stock := repository.FindAll()[0]
				Ω(stock).Should(Equal(weg))
				Ω(mock.ExpectationsWereMet()).Should(BeNil())
			})

		})
	})

	Describe("Getting stock by id", func() {

		Context("When there is no stock", func() {

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "name", "ticker"})
				mock.ExpectQuery("SELECT \\* FROM `stocks` WHERE `stocks`.`id`").WillReturnRows(rows)
			})

			It("should return an error", func() {
				_, err := repository.GetByID("123")
				Ω(err).ShouldNot(BeNil())
				Ω(errors.Is(err, gorm.ErrRecordNotFound)).Should(BeTrue())
			})
		})

		Context("When there is a stock", func() {

			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "name", "ticker"}).AddRow(123, "Weg", "WEGE3")
				mock.ExpectQuery("SELECT \\* FROM `stocks` WHERE `stocks`.`id`").WillReturnRows(rows)
			})

			It("should return the stock", func() {
				stock, err := repository.GetByID("123")
				Ω(stock).Should(Equal(weg))
				Ω(err).Should(BeNil())
			})
		})
	})

	Describe("Saving stocks", func() {

		Context("When add a stock", func() {

			BeforeEach(func() {
				mock.ExpectExec("INSERT INTO `stocks`").WithArgs("Weg", "WEGE3", 123).WillReturnResult(sqlmock.NewResult(123, 1))
			})

			It("should return the saved element", func() {
				stock := repository.Save(weg)
				Ω(stock).Should(Equal(weg))
				Ω(mock.ExpectationsWereMet()).Should(BeNil())
			})
		})
	})

	Describe("Deleting stocks", func() {

		Context("When there is no stock", func() {

			BeforeEach(func() {
				mock.ExpectExec("DELETE FROM `stocks` WHERE `stocks`.`id`").WithArgs("2334").WillReturnResult(sqlmock.NewResult(0, 0))
			})

			It("should return an error", func() {
				err := repository.DeleteByID("2334")
				Ω(err).ShouldNot(BeNil())
				Ω(errors.Is(err, gorm.ErrRecordNotFound)).Should(BeTrue())
			})

		})

		Context("When there is the stock", func() {

			BeforeEach(func() {
				mock.ExpectExec("DELETE FROM `stocks` WHERE `stocks`.`id`").WithArgs("123").WillReturnResult(sqlmock.NewResult(123, 1))
			})

			It("should delete the record", func() {
				err := repository.DeleteByID("123")
				Ω(err).Should(BeNil())
			})

		})

	})

	AfterEach(func() {
		db.Close()
	})
})
