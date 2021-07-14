package siloms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/aplicacao#!/recursos
// https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoDolarDia(dataCotacao=@dataCotacao)?@dataCotacao='06-10-2020'&$top=100&$format=json&$select=cotacaoVenda

// ExtrairCotacao obtém a cotação
func ExtrairFatorCotacao(p Parametro) float64 {
	dtInicial := p.DataInicial.Format("01-02-2006") // MM-DD-AAAA
	dtFinal := p.DataFinal.Format("01-02-2006")

	url := "https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoDolarDia(dataCotacao=@dataCotacao)?@dataCotacao='%s'&$top=100&$format=json&$select=cotacaoVenda"
	urlInicial := fmt.Sprintf(url, dtInicial)
	urlFinal := fmt.Sprintf(url, dtFinal)

	var fator float64
	jsonCotacao := extrairJSON(urlFinal)
	fator = extrairCotacaoFromJSON(jsonCotacao)
	jsonCotacao = extrairJSON(urlInicial)
	fator /= extrairCotacaoFromJSON(jsonCotacao)

	return fator
}

type estrutura struct {
	Value []struct {
		CotacaoVenda float64 `json:"cotacaoVenda"`
	} `json:"value"`
}

func extrairCotacaoFromJSON(jsonData []byte) float64 {
	var e estrutura
	json.Unmarshal(jsonData, &e)
	return e.Value[0].CotacaoVenda
}

func extrairJSON(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}
