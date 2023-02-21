package main

import (
	"github.com/danielcesario/finman/cmd/api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	handler := handler.NewHandler()
	router.POST("/api/transactions", handler.CreateTransaction)

	router.Run("localhost:8080")
}
