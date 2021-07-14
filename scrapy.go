package siloms

import "time"

// Parametro contempla os parâmetros para extração de dados
type Parametro struct {
	DataInicial time.Time
	DataFinal   time.Time
}

/*
// Fator retorna fator de acordo com o parâmetro de entrada
func Fator(p Parametro, f func(Parametro) float64) float64 {
	return f(p)
}
*/
