package main

import (
	"log"

	"github.com/felipemacedo1/go-antifraud-ms/pkg/app/controller"
	"github.com/felipemacedo1/go-antifraud-ms/pkg/app/ml"
	"github.com/felipemacedo1/go-antifraud-ms/pkg/app/service"
	"github.com/gin-gonic/gin"
	"github.com/sjwhitworth/golearn/base"
)

func main() {
	// Load the training dataset
	trainData, err := base.ParseCSVToInstances("path/to/train_data.csv", true)
	if err != nil {
		log.Fatalf("Error loading the training dataset: %v", err)
	}

	model := ml.TrainLinearRegression(trainData)

	fraudService := service.NewCheckFraudService(model)

	fraudController := controller.NewFraudController(fraudService)

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/check-fraud", fraudController.CheckFraudHandler)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}

// func init() {
// 	log.Println("Initializing fraud check service...")

// 	fraudService, err := service.NewCheckFraudService()
// 	if err != nil {
// 		log.Fatalf("Error creating CheckFraudService: %v", err)
// 	}
// 	log.Println("Fraud check service created successfully.")

// 	log.Println("Initializing fraud controller...")
// 	fraudController := controller.NewFraudController(fraudService)
// 	log.Println("Fraud controller initialized successfully.")

// 	log.Println("Setting up server routes...")
// 	router := gin.Default()

// 	// routes
// 	api := router.Group("/api")
// 	{
// 		api.POST("/check-fraud", fraudController.CheckFraudHandler)
// 	}
// 	log.Println("Routes set up successfully.")

// 	log.Println("Starting the server on port 8080...")
// 	// Start server
// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatalf("Error starting the server: %v", err)
// 	}
// 	log.Println("Server started successfully on port 8080.")

// }
