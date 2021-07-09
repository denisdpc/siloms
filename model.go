package siloms

import "time"

// Requisicao contempla campos de interesse da requisicao
type Requisicao struct {
	Numero     string
	PartNumber string
	Status     string
	DataPlano  time.Time
}
