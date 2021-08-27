package siloms

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// MapaPNtoRequisicoes lê arquivo e extrai array de requisições para cada PN
func MapaPNtoRequisicoes(arq string) map[string][]Requisicao {
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

	var mapaRequisicao = make(map[string][]Requisicao) // PN --> requisições

	for scanner.Scan() {
		linha := scanner.Text()
		col := strings.Split(linha, ";")
		numReq := strings.TrimSpace(col[1])
		partNumber := strings.TrimSpace(col[4])
		tipoPlano := strings.TrimSpace(col[14])

		if !isFromParque(numReq) {
			continue
		}

		if !isNacionalizado(partNumber) {
			continue
		}

		if !isMaterial(tipoPlano) { // desconsidera extra-sistema, publicação
			continue
		}

		dataPlano, err := time.Parse("02/01/2006", strings.TrimSpace(col[15][:10]))
		if err != nil {
			log.Fatal(err)
		}

		nomenclatura := strings.TrimSpace(col[7])
		status := strings.TrimSpace(col[17])
		valorUnitStr := strings.TrimSpace(col[28])
		valorUnitStr = strings.Replace(valorUnitStr, ",", ".", 1)
		valorUnit, _ := strconv.ParseFloat(valorUnitStr, 64)
		qtd, _ := strconv.Atoi(strings.TrimSpace(col[30]))
		unidade := strings.TrimSpace(col[31])

		req := Requisicao{
			Numero:       numReq,
			PartNumber:   partNumber,
			Nomenclatura: nomenclatura,
			Status:       status,
			DataPlano:    dataPlano,
			Qtd:          qtd,
			Unidade:      unidade,
			ValorUnit:    valorUnit,
		}

		mapaRequisicao[partNumber] = append(mapaRequisicao[partNumber], req)
	}
	return mapaRequisicao
}

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
		tipoPlano := strings.TrimSpace(col[14])

		if !isFromParque(numReq) {
			continue
		}

		if !isNacionalizado(partNumber) {
			continue
		}

		if !isMaterial(tipoPlano) { // desconsidera extra-sistema, publicação
			continue
		}

		dataPlano, err := time.Parse("02/01/2006", strings.TrimSpace(col[15][:10]))
		if err != nil {
			log.Fatal(err)
		}

		nomenclatura := strings.TrimSpace(col[7])
		status := strings.TrimSpace(col[17])
		valorUnitStr := strings.TrimSpace(col[28])
		valorUnitStr = strings.Replace(valorUnitStr, ",", ".", 1)
		valorUnit, _ := strconv.ParseFloat(valorUnitStr, 64)
		qtd, _ := strconv.Atoi(strings.TrimSpace(col[30]))
		unidade := strings.TrimSpace(col[31])

		req := Requisicao{
			Numero:       numReq,
			PartNumber:   partNumber,
			Nomenclatura: nomenclatura,
			Status:       status,
			DataPlano:    dataPlano,
			Qtd:          qtd,
			Unidade:      unidade,
			ValorUnit:    valorUnit,
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

// isMaterial verifica se a requisição não é extra-sistema, publicação
func isMaterial(tipoPlano string) bool {
	return tipoPlano == "Material"
}
