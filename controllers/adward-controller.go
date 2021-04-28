package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hieudt-2054/report-adwards/models"
	"github.com/hieudt-2054/report-adwards/services"
)

type AdwardController interface {
	All(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type controller struct {
	service services.AdwardService
}

func New(service services.AdwardService) AdwardController {
	return &controller{
		service: service,
	}
}

func (c *controller) All(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.All())
}

func (c *controller) Save(ctx *gin.Context) {
	var adward models.Adward
	ctx.ShouldBindJSON(&adward)
	c.service.Save(adward)
	ctx.JSON(http.StatusCreated, adward)
}
