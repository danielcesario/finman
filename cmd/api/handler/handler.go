package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionService interface{}

type Handler struct {
	Service TransactionService
}

func NewHandler(service TransactionService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateTransaction(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusCreated, map[string]interface{}{"result": "ok"})
}
