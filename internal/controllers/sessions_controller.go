package controllers

import (
	"net/http"
	"sps-backend/internal/services"
	"sps-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type SessionsController struct {
	sessionLogger  *utils.Logger
	jwtSecret      string
	sessionService *services.SessionServices
}

func NewSessionController(jwtSecret string, sessionServices *services.SessionServices) *SessionsController {
	return &SessionsController{
		jwtSecret:      jwtSecret,
		sessionLogger:  utils.GetUserAuthLogger(),
		sessionService: sessionServices,
	}
}

func (c *SessionsController) GetToken(ctx *gin.Context) {

	ipAddress := ctx.GetHeader("X-IP-ADDRESS")
	timestamp := ctx.GetHeader("X-TIMESTAMP")

	// terserah mau nambah apa lagi

	res, err := c.sessionService.GetTokenService(ctx, ipAddress, timestamp)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.SuccessUnwrapped(ctx, http.StatusOK, res)

}
