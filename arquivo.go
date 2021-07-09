package siloms

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// LerArqRequisicao extrai as requisições de um arquivo no formato CSV
func LerArqRequisicao(arq string) []Requisicao {
	file, err := os.Open(arq)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dec := transform.NewReader(file, charmap.ISO8859_1.NewDecoder())
	scanner := bufio.NewScanner(dec)
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()

	var requisicoes []Requisicao

	for scanner.Scan() {
		linha := scanner.Text()
		col := strings.Split(linha, ";")
		numReq := strings.TrimSpace(col[1])
		partNumber := strings.TrimSpace(col[4])

		if !isFromParque(numReq) {
			continue
		}

		if !isNacionalizado(partNumber) {
			continue
		}

		dataPlano, _ := time.Parse("02/01/2006", strings.TrimSpace(col[15][:10]))

		req := Requisicao{
			Numero:     numReq,
			PartNumber: partNumber,
			Status:     strings.TrimSpace(col[17]),
			DataPlano:  dataPlano,
		}

		requisicoes = append(requisicoes, req)

	}

	return requisicoes
}

// isFromParque verifica se a requisição foi emitida por algum Parque de Material
func isFromParque(numReq string) bool {
	return strings.HasPrefix(numReq, "GL") ||
		strings.HasPrefix(numReq, "LS") ||
		strings.HasPrefix(numReq, "SP") ||
		strings.HasPrefix(numReq, "PB")
}

// isNacionalizado verifica se a requisição é de item nacionalizado
func isNacionalizado(partNumber string) bool {
	return strings.HasPrefix(partNumber, "DCN")
}

// TODO
// isExtraSistema verifica se a requisição é extra sistema
func isExtraSistema(es string) bool {
	return es == "Sim"
}
