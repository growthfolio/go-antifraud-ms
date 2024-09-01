package ml

import (
	"fmt"
	"strconv"

	"github.com/felipemacedo1/go-antifraud-ms/pkg/app/models"
	"github.com/sjwhitworth/golearn/base"
)

func TransactionToInstances(transaction models.Transaction) (base.FixedDataGrid, error) {
	data := base.NewDenseInstances()

	attributeNames := []string{
		"idade",
		"renda_anual",
		"historico_credito",
		"valor",
		"tipo_transacao",
		"local_transacao",
		"categoria_comercio",
		"canal_autenticacao",
		"frequencia_transacoes_24h",
		"distancia_localizacao",
		"tentativas_falhas_ultimas_24h",
		"hora_transacao",
	}

	valueMap := map[string]interface{}{
		"idade":                         transaction.Idade,
		"renda_anual":                   transaction.RendaAnual,
		"historico_credito":             transaction.HistoricoCredito,
		"valor":                         transaction.Valor,
		"tipo_transacao":                transaction.TipoTransacao,
		"local_transacao":               transaction.LocalTransacao,
		"categoria_comercio":            transaction.CategoriaComercio,
		"canal_autenticacao":            transaction.CanalAutenticacao,
		"frequencia_transacoes_24h":     transaction.FrequenciaTransacoes24h,
		"distancia_localizacao":         transaction.DistanciaLocalizacao,
		"tentativas_falhas_ultimas_24h": transaction.TentativasFalhasUltimas24h,
		"hora_transacao":                transaction.HoraTransacao,
	}

	// Define the attributes and add them to the dataset
	attrs := make([]base.Attribute, len(attributeNames))
	for i, name := range attributeNames {
		attrs[i] = base.NewFloatAttribute(name)
		data.AddAttribute(attrs[i])
	}
	data.AddClassAttribute(attrs[len(attrs)-1])

	// Initialize the attributeSpecs map
	attributeSpecs := make(map[string]base.AttributeSpec)

	// Retrieve attribute specifications
	for i, attr := range data.AllAttributes() {
		name := attributeNames[i]
		if attr.GetName() == name {
			attrSpec, err := data.GetAttribute(attr)
			if err != nil {
				return nil, fmt.Errorf("failed to get attribute spec for %s: %v", name, err)
			}
			attributeSpecs[name] = attrSpec
		}
	}

	// Extend the dataset to accommodate a single instance
	data.Extend(1)

	// Set values in the data instance
	for name, value := range valueMap {
		spec, exists := attributeSpecs[name]
		if !exists {
			return nil, fmt.Errorf("attribute %s not found in attributeSpecs", name)
		}

		// Perform type conversion based on the expected attribute type
		switch v := value.(type) {
		case int:
			data.Set(spec, 0, []byte(strconv.Itoa(v)))
		case float64:
			data.Set(spec, 0, []byte(strconv.FormatFloat(v, 'f', -1, 64)))
		case string:
			data.Set(spec, 0, []byte(v))
		default:
			return nil, fmt.Errorf("unsupported attribute type for %s: %T", name, value)
		}
	}

	return data, nil
}
