package siloms

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

// Requisicao contempla campos de interesse da requisicao
type Requisicao struct {
	numero     string
	partNumber string
	status     string
}

func Teste() {
	fmt.Println("teste req")
}

// VerificarReqNacionalizado verifica se a requisição é de material nacionalizado
func VerificarReqNacionalizado(r Requisicao) bool {
	return strings.HasPrefix(r.partNumber, "DCN")
}

// VerificarReqPendente verifica se a requisição não atingiu o status de mapa
func VerificarReqPendente(r Requisicao) bool {

	return false
}

// VerificarReqAtendida verifica se a requisição atingiu o status de mapa e posteriores
func VerificarReqAtendida(r Requisicao) bool {
	return false
}

// ExtrairRequisicoes extrai requisições que atende a determinado requisito
func ExtrairRequisicoes(v func(Requisicao) bool) []Requisicao {
	return nil
}

// LerArqRequisicao extrai as requisições de um arquivo no formato CSV
func LerArqRequisicao(arq string) []Requisicao {

	reader := csv.NewReader(bytes.NewBuffer(arq))
	_, err := reader.Read() // skip first line
	if err != nil {
		if err != io.EOF {
			log.Fatalln(err)
		}
	}
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				break
			}
		}
		fmt.Println(line)
	}

	return nil
}
