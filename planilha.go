// Package siloms inclui funções para composição de planilha
package siloms

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"time"
)

// Planilhar gerar planilha a partir de parâmetros de entrada
func Planilhar(caminho string, dataPesq time.Time,
	reqRef Requisicao, reqs []Requisicao,
	pTDPre, pCot, pIGPM, pEsc float64) {

	p := getPlanilhaModelo(caminho)
	if p == nil {
		return
	}

	setarValores(p, dataPesq, reqRef, reqs, pTDPre, pCot, pIGPM, pEsc)
	setarFormulas(p)

	if err := p.SaveAs(caminho + "teste.xlsx"); err != nil {
		fmt.Println("ERRO: ", err)
	}

}

func setarValores(p *excelize.File, dataPesq time.Time,
	reqRef Requisicao, reqs []Requisicao,
	pTDPre, pCot, pIGPM, pEsc float64) {

	p.SetCellValue("planilha", "E6", reqRef.PartNumber)
	p.SetCellValue("planilha", "E7", reqRef.Nomenclatura)
	p.SetCellValue("planilha", "A11", reqRef.DataPlano.Format("02/01/2006"))
	p.SetCellValue("planilha", "B11", reqRef.Numero)
	p.SetCellValue("planilha", "G11", reqRef.Qtd)
	p.SetCellValue("planilha", "H11", reqRef.Unidade)
	p.SetCellValue("planilha", "I11", reqRef.ValorUnit)

	p.SetCellValue("planilha", "J15", pTDPre)
	p.SetCellValue("planilha", "J16", pEsc)
	p.SetCellValue("planilha", "J17", pCot)
	p.SetCellValue("planilha", "J18", pIGPM)
	p.SetCellValue("planilha", "J19", dataPesq.Format("02/01/2006"))

	var unidade string

	mesmaUnidade := true
	qtdTotal := 0
	for i, req := range reqs {
		linha := strconv.Itoa(15 + i)
		p.SetCellValue("planilha", "A"+linha, req.Numero)
		p.SetCellValue("planilha", "E"+linha, req.Qtd)
		p.SetCellValue("planilha", "F"+linha, req.Unidade)

		if unidade == "" {
			unidade = req.Unidade
		}

		mesmaUnidade = mesmaUnidade && (unidade == req.Unidade)
		qtdTotal += req.Qtd
	}
	if !mesmaUnidade {
		qtdTotal = -1
		unidade = "---"
	}
	p.SetCellValue("planilha", "E18", qtdTotal)
	p.SetCellValue("planilha", "F18", unidade)

}

func setarFormulas(p *excelize.File) {
	p.SetCellFormula("planilha", "A27", "G11*I11/(1+G11*J16)*J18*(1+J15)")
	p.SetCellFormula("planilha", "C27", "G11*I11/(1+G11*J16)*J16*J17*(1+J15)")
	p.SetCellFormula("planilha", "E27", "A27/E18+C27")
	p.SetCellFormula("planilha", "H27", "E26*E18")
}

func getPlanilhaModelo(caminho string) *excelize.File {
	f, err := excelize.OpenFile(caminho + "modelo.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return f
}
