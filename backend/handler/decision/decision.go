package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gabrielvinte/PergunteAoBalancaRaboLeitoso/handler"
	decisionRequestDto "github.com/gabrielvinte/PergunteAoBalancaRaboLeitoso/handler/decision/decisionDto"
	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TakeDecisionHandler(ctx *gin.Context) {
	request := decisionRequestDto.DecisionRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	if request.OptionA == "" || request.OptionB == "" {
		handler.SendError(ctx, http.StatusBadRequest, "Invalid option A or B")
		return
	}
	response := randomizeOptions(request.OptionA, request.OptionB)
	handler.SendSucess(ctx, "take decision", response)
}

func randomizeOptions(optionA string, optionB string) string {
	options := []string{optionA, optionB}
	return options[rand.Intn(len(options))]
}