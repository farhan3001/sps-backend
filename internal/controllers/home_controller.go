package controllers

import (
	"net/http"
	"sps-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func NewHomeController(homeServices string) *HomeController {
	return &HomeController{}
}

func (c *HomeController) Home(ctx *gin.Context) {
	utils.Success(ctx, http.StatusOK, "Ok")
}

func (c *HomeController) Health(ctx *gin.Context) {
	utils.Success(ctx, http.StatusOK, "Health Ok")
}
