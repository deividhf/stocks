package main

import (
	"log"

	"github.com/deividhf/stocks/config"
	"github.com/deividhf/stocks/stocks"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var err error
	dns := "root:root@tcp(mysql:3306)/stocks?charset=utf8&parseTime=True&loc=Local"
	config.DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error opening database. %s", err)
	}

	setupServer().Run(":8080")
}

func setupServer() *gin.Engine {
	server := gin.Default()

	router := stocks.DefaultRouter()
	router.Routes(server)

	return server
}
