package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateTransaction(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusCreated, map[string]interface{}{"result": "ok"})
}
