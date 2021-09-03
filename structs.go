package siloms

import (
	"strconv"
	"strings"
	"time"
)

// Parametro contempla os parâmetros para extração de dados
type Parametro struct {
	DataInicial time.Time
	DataFinal   time.Time
}

// Requisicao contempla campos de interesse da requisicao
type Requisicao struct {
	Numero       string
	PartNumber   string
	Nomenclatura string
	Status       string
	DataPlano    time.Time
	Qtd          int
	Unidade      string
	ValorUnit    float64
}

// RequisicaoAtualRef relação de requisições atuais e de
// referência para um dado part number
type RequisicaoPnRefToAtual struct {
	PartNumber string
	ReqsRef    []Requisicao
	ReqsAtual  []Requisicao
}

// FormatarData retorna data no formato AAAA-MM-DD
func (r Requisicao) FormatarAno() string {
	return r.DataPlano.Format("2006")
	//return r.DataPlano.Format("2006-01-02")
}

// FormatarValor converte ValorUnit para string e substitui . por ,
func (r Requisicao) FormatarValor() string {
	str := strconv.FormatFloat(r.ValorUnit, 'f', 2, 64)
	return strings.Replace(str, ".", ",", 1)
}
