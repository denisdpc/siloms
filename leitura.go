package siloms

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

func Teste() {
	fmt.Println("teste siloms")
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
