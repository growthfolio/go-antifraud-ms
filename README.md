# 🛡️ Go Antifraud Microservice - Detecção de Fraudes

## 🎯 Objetivo de Aprendizado
Microserviço desenvolvido em Go para estudar **detecção de fraudes** em tempo real e **machine learning aplicado**. Implementa algoritmos de análise comportamental e padrões suspeitos em transações financeiras, aplicando **Clean Architecture** e **observabilidade**.

## 🛠️ Tecnologias Utilizadas
- **Linguagem:** Go
- **Arquitetura:** Clean Architecture, Microservices
- **Machine Learning:** Algoritmos de detecção de anomalias
- **Monitoramento:** Logging estruturado, métricas
- **Containerização:** Docker
- **Banco de dados:** PostgreSQL, Redis (cache)
- **Comunicação:** REST API, gRPC

## 🚀 Demonstração
```go
// Análise de transação em tempo real
type TransactionAnalysis struct {
    TransactionID string    `json:"transaction_id"`
    Amount        float64   `json:"amount"`
    RiskScore     float64   `json:"risk_score"`
    IsFraud       bool      `json:"is_fraud"`
    Reasons       []string  `json:"reasons"`
}

// POST /api/v1/transactions/analyze
{
  "transaction_id": "tx_123456",
  "amount": 5000.00,
  "merchant": "Online Store",
  "location": "BR",
  "timestamp": "2024-01-15T10:30:00Z"
}
```

## 📁 Estrutura do Projeto
```
go-antifraud-ms/
├── cmd/
│   └── main.go                # Entry point
├── pkg/
│   ├── antifraud/            # Core fraud detection
│   │   ├── analyzer.go       # Analysis engine
│   │   ├── rules.go          # Business rules
│   │   └── ml_models.go      # ML algorithms
│   ├── api/                  # HTTP handlers
│   ├── config/               # Configuration
│   ├── database/             # Data persistence
│   ├── logger/               # Structured logging
│   └── server/               # Server setup
├── scripts/                  # Deployment scripts
└── docker-compose.yml        # Local development
```

## 💡 Principais Aprendizados

### 🔍 Detecção de Fraudes
- **Análise comportamental:** Padrões de usuário anômalos
- **Regras de negócio:** Limites e validações automáticas
- **Machine learning:** Algoritmos de classificação
- **Real-time processing:** Análise em tempo real
- **Risk scoring:** Sistema de pontuação de risco

### 🏗️ Arquitetura de Microserviços
- **Single responsibility:** Foco específico em antifraude
- **Scalability:** Capacidade de escalar independentemente
- **Resilience:** Tolerância a falhas e degradação graceful
- **Observability:** Logs, métricas e tracing distribuído
- **API design:** Interfaces bem definidas

### 🤖 Machine Learning Integration
- **Feature engineering:** Extração de características relevantes
- **Model training:** Treinamento com dados históricos
- **Real-time inference:** Predições em tempo real
- **Model versioning:** Versionamento de modelos
- **Performance monitoring:** Monitoramento de acurácia

## 🧠 Conceitos Técnicos Estudados

### 1. **Fraud Detection Engine**
```go
type FraudAnalyzer struct {
    ruleEngine    RuleEngine
    mlModel       MLModel
    riskThreshold float64
}

func (fa *FraudAnalyzer) AnalyzeTransaction(tx Transaction) (*Analysis, error) {
    // Rule-based analysis
    ruleScore := fa.ruleEngine.Evaluate(tx)
    
    // ML-based analysis
    mlScore := fa.mlModel.Predict(tx.Features())
    
    // Combined risk score
    finalScore := (ruleScore + mlScore) / 2
    
    return &Analysis{
        RiskScore: finalScore,
        IsFraud:   finalScore > fa.riskThreshold,
        Reasons:   fa.generateReasons(tx, ruleScore, mlScore),
    }, nil
}
```

### 2. **Business Rules Engine**
```go
type Rule interface {
    Evaluate(tx Transaction) (score float64, reason string)
}

type AmountRule struct {
    MaxAmount float64
}

func (r *AmountRule) Evaluate(tx Transaction) (float64, string) {
    if tx.Amount > r.MaxAmount {
        return 0.8, "Transaction amount exceeds limit"
    }
    return 0.0, ""
}
```

### 3. **Real-time Processing**
```go
func (s *Server) handleTransactionAnalysis(w http.ResponseWriter, r *http.Request) {
    var tx Transaction
    if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    // Async analysis for performance
    analysis, err := s.analyzer.AnalyzeTransaction(tx)
    if err != nil {
        s.logger.Error("Analysis failed", "error", err)
        http.Error(w, "Analysis failed", http.StatusInternalServerError)
        return
    }
    
    // Log for audit trail
    s.logger.Info("Transaction analyzed", 
        "tx_id", tx.ID, 
        "risk_score", analysis.RiskScore,
        "is_fraud", analysis.IsFraud)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(analysis)
}
```

## 🚧 Desafios Enfrentados
1. **False positives:** Balanceamento entre segurança e usabilidade
2. **Real-time performance:** Latência baixa em análises complexas
3. **Model accuracy:** Melhoria contínua dos algoritmos
4. **Data quality:** Tratamento de dados inconsistentes
5. **Scalability:** Handling de alto volume de transações

## 📚 Recursos Utilizados
- [Fraud Detection Techniques](https://www.sciencedirect.com/topics/computer-science/fraud-detection)
- [Machine Learning for Fraud Detection](https://towardsdatascience.com/machine-learning-for-fraud-detection-3b8b8b8b8b8b)
- [Microservices Patterns](https://microservices.io/patterns/)
- [Go Best Practices](https://golang.org/doc/effective_go.html)

## 📈 Próximos Passos
- [ ] Implementar deep learning models
- [ ] Adicionar análise de grafos de relacionamentos
- [ ] Criar dashboard de monitoramento
- [ ] Implementar feedback loop para ML
- [ ] Adicionar análise de comportamento temporal
- [ ] Integrar com sistemas de pagamento

## 🔗 Projetos Relacionados
- [AMQP Transactions MS](../amqp-transactions-ms/) - Processamento de transações
- [Go PriceGuard API](../go-priceguard-api/) - API Go com Clean Architecture
- [CryptoTool](../CryptoTool/) - Segurança e criptografia

---

**Desenvolvido por:** Felipe Macedo  
**Contato:** contato.dev.macedo@gmail.com  
**GitHub:** [FelipeMacedo](https://github.com/felipemacedo1)  
**LinkedIn:** [felipemacedo1](https://linkedin.com/in/felipemacedo1)

> 💡 **Reflexão:** Este projeto proporcionou compreensão profunda sobre segurança financeira e machine learning aplicado. A implementação de sistemas de detecção de fraudes demonstrou a importância da análise em tempo real e da arquitetura resiliente.