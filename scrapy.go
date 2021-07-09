package siloms

import (
	"fmt"
	"github.com/gocolly/colly"
)

// Parametro contempla os parâmetros para extração de dados
type Parametro struct {
	dataInicial string
	dataAtual   string
}

func Fator(p Parametro, f func(Parametro) float64) float64 {
	return f(p)
}

func ExtrairCotacao(p Parametro) float64 {
	c := colly.NewCollector()
	fmt.Println(c)
	return 0.0
}

func ExtrairIGPM(p Parametro) float64 {
	return 0.0
}

func ExtrairTaxaTesouro(p Parametro) float64 {
	return 0.0
}
