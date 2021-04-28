package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hieudt-2054/report-adwards/controllers"
	services "github.com/hieudt-2054/report-adwards/services"
)

var (
	adwardService    services.AdwardService       = services.New()
	adwardController controllers.AdwardController = controllers.New(adwardService)
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()
	v1 := routes.Group("/api/v1")
	{
		v1.GET("adward", adwardController.All)
		v1.POST("adward", adwardController.Save)
		v1.GET("/healthcheck", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
		})
	}

	return routes
}
