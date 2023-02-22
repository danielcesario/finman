package handler

import (
	"context"
	"net/http"

	"github.com/danielcesario/finman/internal/transactions"
	"github.com/gin-gonic/gin"
)

type TransactionService interface {
	CreateUser(ctx context.Context, request transactions.UserRequest) (*transactions.UserResponse, error)
}

type Handler struct {
	Service TransactionService
}

func NewHandler(service TransactionService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Register(context *gin.Context) {
	var request transactions.UserRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := request.HashPassword(request.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	response, err := h.Service.CreateUser(context, request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, response)
}
