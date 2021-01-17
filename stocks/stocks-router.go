package stocks

import (
	"net/http"

	"github.com/deividhf/stocks/stocks/controller"
	"github.com/deividhf/stocks/stocks/service"
	"github.com/gin-gonic/gin"
)

var (
	s = service.New()
	c = controller.New(s)
)

// Routes maps all routes from stocks
func Routes(route *gin.Engine) {
	stocks := route.Group("/stocks")

	stocks.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, c.FindAll())
	})

	stocks.POST("", func(ctx *gin.Context) {
		stock, err := c.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusCreated, stock)
		}
	})

}
