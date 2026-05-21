package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"message":     message,
		"status_code": statusCode,
	})
}
func SendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "aplication/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s sucessfull", op),
		"data":    data,
	})
}