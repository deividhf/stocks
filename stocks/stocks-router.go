package stocks

import (
	"net/http"

	"github.com/deividhf/stocks/config"
	"github.com/deividhf/stocks/stocks/controller"
	"github.com/deividhf/stocks/stocks/repository"
	"github.com/gin-gonic/gin"
)

// StockRouter is the router of stocks
type StockRouter interface {
	Routes(route *gin.Engine)
}

type stockRouter struct {
	controller controller.StockController
}

// DefaultRouter returns the default router
func DefaultRouter() StockRouter {
	return &stockRouter{
		controller: controller.New(repository.New(config.DB)),
	}
}

// New creates the a new router receiving a controller
func New(controller controller.StockController) StockRouter {
	return &stockRouter{
		controller: controller,
	}
}

// Routes maps all routes from stocks
func (r *stockRouter) Routes(route *gin.Engine) {
	stocks := route.Group("/stocks")

	stocks.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, r.controller.FindAll())
	})

	stocks.GET("/:stock_id", func(ctx *gin.Context) {
		stock, err := r.controller.GetByID(ctx.Param("stock_id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, stock)
		}
	})

	stocks.POST("", func(ctx *gin.Context) {
		stock, err := r.controller.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusCreated, stock)
		}
	})

}
