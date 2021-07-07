package siloms

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// Requisicao contempla campos de interesse da requisicao
type Requisicao struct {
	numero     string
	partNumber string
	status     string
}

// Teste função de teste
func Teste() {
	fmt.Println("siloms TESTE2")
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
		req := Requisicao{
			numero:     strings.TrimSpace(col[1]),
			partNumber: strings.TrimSpace(col[4]),
			status:     strings.TrimSpace(col[17])}

		requisicoes = append(requisicoes, req)
	}

	return requisicoes
}
