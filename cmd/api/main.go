package main

import (
	"github.com/danielcesario/finman/cmd/api/handler"
	"github.com/danielcesario/finman/internal/database"
	"github.com/danielcesario/finman/internal/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.OpenDb()
	if err != nil {
		panic("failed to connect database")
	}

	database.UpdateSchema(db)

	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	handler := handler.NewHandler(service)

	router := gin.Default()
	router.POST("/api/register", handler.Register)

	router.Run("localhost:8080")
}
