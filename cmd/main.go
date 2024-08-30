package main

import (
	"log"

	"github.com/felipemacedo1/go-antifraud-ms/services/internal/app/controller/"
	"github.com/felipemacedo1/go-antifraud-ms/services/internal/app/service/"
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
