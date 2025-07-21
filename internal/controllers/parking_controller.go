package controllers

import (
	"net/http"
	"sps-backend/internal/domain"
	"sps-backend/internal/services"
	"sps-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type ParkingController struct {
	parkingServices *services.ParkingServices
}

func NewParkingController(parkingServices *services.ParkingServices) *ParkingController {
	return &ParkingController{
		parkingServices: parkingServices,
	}
}

func (c *ParkingController) ParkingInquiry(ctx *gin.Context) {

	var req domain.ParkingInquiryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := c.parkingServices.ParkingInq(ctx, &req)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	utils.SuccessUnwrapped(ctx, http.StatusOK, res)
}
