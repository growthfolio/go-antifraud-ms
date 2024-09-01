package ml

import (
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/linear_models"
)

// Model encapsulates the trained linear regression model
type Model struct {
	Linear *linear_models.LinearRegression
}

// TrainLinearRegression trains a linear regression model
func TrainLinearRegression(trainData base.FixedDataGrid) *Model {
	linear := linear_models.NewLinearRegression()

	err := linear.Fit(trainData)
	if err != nil {
		log.Fatalf("Error training the linear regression model: %v", err)
	}

	return &Model{Linear: linear}
}

// Predict makes predictions using the trained model
func (m *Model) Predict(data base.FixedDataGrid) (base.FixedDataGrid, error) {
	return m.Linear.Predict(data)
}
