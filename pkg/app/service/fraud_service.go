package service

import (
	"github.com/felipemacedo1/go-antifraud-ms/pkg/app/models"

	"github.com/felipemacedo1/go-antifraud-ms/pkg/app/ml"
)

type CheckFraudService struct {
	model *ml.Model
}

func NewCheckFraudService(model *ml.Model) *CheckFraudService {
	return &CheckFraudService{model: model}
}

func (s *CheckFraudService) Execute(transaction models.Transaction) (string, error) {
	// Converter a transação para um formato que GoLearn entende
	data, err := ml.TransactionToInstances(transaction)
	if err != nil {
		return "", err
	}

	// Fazer a previsão usando o modelo treinado
	predictions, err := s.model.Predict(data)
	if err != nil {
		return "", err
	}

	// Avaliar e retornar o resultado
	result := predictions.RowString(0)
	return result, nil
}
