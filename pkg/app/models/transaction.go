package models

type Transaction struct {
	Idade                      int     `json:"idade"`
	RendaAnual                 float64 `json:"renda_anual"`
	HistoricoCredito           float64 `json:"historico_credito"`
	Valor                      float64 `json:"valor"`
	TipoTransacao              string  `json:"tipo_transacao"`
	LocalTransacao             string  `json:"local_transacao"`
	CategoriaComercio          string  `json:"categoria_comercio"`
	CanalAutenticacao          string  `json:"canal_autenticacao"`
	FrequenciaTransacoes24h    int     `json:"frequencia_transacoes_24h"`
	DistanciaLocalizacao       float64 `json:"distancia_localizacao"`
	TentativasFalhasUltimas24h int     `json:"tentativas_falhas_ultimas_24h"`
	HoraTransacao              string  `json:"hora_transacao"`
}
