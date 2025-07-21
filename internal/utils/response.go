package utils

import (
	"github.com/gin-gonic/gin"
)

var (
	ErrUnauthorized = gin.H{"error": "unauthorized"}
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, Response{
		Data: data,
	})
}

func SuccessUnwrapped(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}

func Error(ctx *gin.Context, status int, err interface{}) {
	ctx.JSON(status, Response{
		Error: err,
	})
}
