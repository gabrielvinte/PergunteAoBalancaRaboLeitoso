package router

import (
	handler "github.com/gabrielvinte/PergunteAoBalancaRaboLeitoso/handler/decision"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	basePath := "/api/"

	v1 := router.Group(basePath)
	{
		v1.POST("/", handler.TakeDecisionHandler)

		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}
}