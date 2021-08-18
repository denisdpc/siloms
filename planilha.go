// Package siloms inclui funções para composição de planilha
package siloms

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"time"
)

// Planilhar gera planilha a partir de parâmetros de entrada
func Planilhar(caminho string, dataPesq time.Time,
	reqRef Requisicao, reqs []Requisicao,
	pTDPre, pCot, pIGPM, pR float64) {
	qtdTotal := 0
	for _, req := range reqs {
		qtdTotal += req.Qtd
	}
	qtdRef, valorUnitRef, dataRef := reqRef.Qtd, reqRef.ValorUnit, reqRef.DataPlano
	fmt.Println(qtdTotal, qtdRef, valorUnitRef, dataRef)

	p := getPlanilhaModelo(caminho)
	if p == nil {
		return
	}
	p.SetCellValue("planilha", "E6", reqRef.PartNumber)
	p.SetCellValue("planilha", "E7", reqRef.Nomenclatura)
	p.SetCellValue("planilha", "A11", reqRef.DataPlano.Format("02/01/2006"))
	p.SetCellValue("planilha", "B11", reqRef.Numero)
	p.SetCellValue("planilha", "G11", reqRef.Qtd)
	p.SetCellValue("planilha", "H11", reqRef.Unidade)
	p.SetCellValue("planilha", "I11", reqRef.ValorUnit)

	if err := p.SaveAs(caminho + "teste.xlsx"); err != nil {
		fmt.Println("ERRO: ", err)
	}

	//fmt.Println(p)

	// f, err := excelize.OpenFile(caminho + "modelo.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(f)
}

func getPlanilhaModelo(caminho string) *excelize.File {
	f, err := excelize.OpenFile(caminho + "modelo.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return f
}
