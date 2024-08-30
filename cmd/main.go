// cmd/main.go
package main

import (
	"go-antifraud-ms/services/internal/app/controller"
	"go-antifraud-ms/services/internal/app/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	fraudService, err := service.NewCheckFraudService()
	if err != nil {
		log.Fatalf("Erro ao inicializar o serviço de fraude: %v", err)
	}

	fraudController := controller.NewFraudController(fraudService)

	router := gin.Default()

	// routes
	api := router.Group("/api")
	{
		api.POST("/check-fraud", fraudController.CheckFraudHandler)
	}

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
