// Package siloms inclui funções para composição de planilha
package siloms

import (
	"fmt"
	"time"
)

// Planilhar gera planilha a partir de parâmetros de entrada
func Planilhar(dataPesq time.Time,
	reqRef Requisicao, reqs []Requisicao,
	pTDPre, pCot, pIGPM, pR float64) {
	qtdTotal := 0
	for _, req := range reqs {
		qtdTotal += req.Qtd
	}
	qtdRef, valorUnitRef, dataRef := reqRef.Qtd, reqRef.ValorUnit, reqRef.DataPlano
	fmt.Println(qtdTotal, qtdRef, valorUnitRef, dataRef)

}
