package routes

import (
	"net/http"
	"strconv"

	"github.com/AvengersCodeLovers/report-adwards/middleware"
	"github.com/AvengersCodeLovers/report-adwards/util"

	controllers "github.com/AvengersCodeLovers/report-adwards/controllers"
	services "github.com/AvengersCodeLovers/report-adwards/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	adwardService    services.AdwardService       = services.New()
	adwardController controllers.AdwardController = controllers.New(adwardService)
)

func SetupRouter() *gin.Engine {
	routes := gin.New()
	routes.Use(middleware.RequestLogMiddleware())
	allowKeepRequestId, _ := strconv.ParseBool(util.GetEnv("REQUEST_ID", "false"))
	routes.Use(middleware.RequestID(allowKeepRequestId))
	routes.Use(gin.Recovery())

	gin.SetMode(util.GetEnv("APP_ENV", "debug"))
	logrus.Infof("Application running in port: %s", util.GetEnv("APP_PORT", "8888"))

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
