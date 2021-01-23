package main

import (
	"log"

	"github.com/deividhf/stocks/config"
	"github.com/deividhf/stocks/stocks"
	"github.com/deividhf/stocks/stocks/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var err error
	config.DB, err = gorm.Open(sqlite.Open("stocks.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error opening database. %s", err)
	}
	config.DB.AutoMigrate(&entity.Stock{})

	setupServer().Run(":8090")
}

func setupServer() *gin.Engine {
	server := gin.Default()

	router := stocks.DefaultRouter()
	router.Routes(server)

	return server
}
