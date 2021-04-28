package routes

import (
	controllers "github.com/AvengersCodeLovers/report-adwards/controllers"
	services "github.com/AvengersCodeLovers/report-adwards/services"
	"github.com/gin-gonic/gin"
)

var (
	adwardService    services.AdwardService       = services.New()
	adwardController controllers.AdwardController = controllers.New(adwardService)
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()
	v1 := routes.Group("/api/v1")
	{
		v1.GET("adward", adwardController.Index)
		v1.POST("adward", adwardController.Store)
		v1.GET("/healthcheck", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}

	return routes
}
