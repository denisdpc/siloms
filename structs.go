package siloms

import "time"

// Parametro contempla os parâmetros para extração de dados
type Parametro struct {
	DataInicial time.Time
	DataFinal   time.Time
}

// Requisicao contempla campos de interesse da requisicao
type Requisicao struct {
	Numero     string
	PartNumber string
	Status     string
	DataPlano  time.Time
}
