package routes

import (
	"log"
	"net/http"
	"os"

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
	gin.SetMode(os.Getenv("APP_ENV"))
	log.Printf("Application running in : %v", gin.Mode())
	v1 := routes.Group("/api/v1")
	{
		v1.GET("adward", adwardController.Index)
		v1.POST("adward", adwardController.Store)
		v1.GET("/healthcheck", healthCheck)
	}

	return routes
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
