package service

import (
	"go-antifraud-ms/services/internal/app/models"
	"go-antifraud-ms/utils"
	"strconv"
)

// CheckFraudService define o serviço de verificação de fraude
type CheckFraudService struct {
	// Por exemplo:
	// model        *Model
	// preprocessor *PreprocessorType
}

// NewCheckFraudService cria uma nova instância do serviço de verificação de fraude
func NewCheckFraudService() (*CheckFraudService, error) {
	service := &CheckFraudService{}
	// Carregue o modelo e o preprocessor aqui
	// Exemplo:
	// model, err := loadModel("path/to/fraud_detection_model.pkl")
	// if err != nil {
	//     return nil, err
	// }
	// preprocessor, err := loadPreprocessor("path/to/preprocessor.pkl")
	// if err != nil {
	//     return nil, err
	// }
	// service.model = model
	// service.preprocessor = preprocessor

	return service, nil
}

func (s *CheckFraudService) Execute(transaction models.Transaction) (string, float64, error) {
	minutes, err := utils.HourToMinutes(transaction.HoraTransacao)
	if err != nil {
		return "", 0, err
	}
	transaction.HoraTransacao = strconv.Itoa(minutes)

	// Pré-processamento da transação:
	// preprocessedData, err := s.preprocessor.Transform(transaction)
	// if err != nil {
	//     return "", 0, err
	// }

	// Exemplo:
	// prediction, proba, err := s.model.Predict(preprocessedData)
	// if err != nil {
	//     return "", 0, err
	// }

	// result := "legítima"
	// if prediction == 1 {
	//     result = "fraudulenta"
	// }

	// mocked values
	result := "legítima"
	proba := 0.95

	return result, proba, nil
}
