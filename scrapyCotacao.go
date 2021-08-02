package siloms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/aplicacao#!/recursos
// https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoDolarDia(dataCotacao=@dataCotacao)?@dataCotacao='06-10-2020'&$top=100&$format=json&$select=cotacaoVenda

// ExtrairCotacao fator de cotação, data inicial e data final
// Há datas que não contemplam cotação, sendo pesquisadas datas anteriores
// As datas efetivamente utilizadas na cotação são fornecidas
func ExtrairFatorCotacao(p *Parametro) (float64, time.Time, time.Time) {
	cotacaoInicial, dataInicial := lerCotacaoDataValida(p.DataInicial)
	cotacaoFinal, dataFinal := lerCotacaoDataValida(p.DataFinal)
	fator := cotacaoFinal / cotacaoInicial
	return fator, dataInicial, dataFinal
}

func lerCotacaoDataValida(dataCotacao time.Time) (float64, time.Time) {
	cotacao := -1.0
	for cotacao == -1 {
		cotacao = lerCotacaoData(dataCotacao)
		if cotacao == -1 {
			dataCotacao = dataCotacao.Add(-time.Hour * 24)
		}
	}
	return cotacao, dataCotacao
}

func lerCotacaoData(date time.Time) float64 {
	dt := date.Format("01-02-2006")
	url := "https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoDolarDia(dataCotacao=@dataCotacao)?@dataCotacao='%s'&$top=100&$format=json&$select=cotacaoVenda"
	urlDate := fmt.Sprintf(url, dt)
	jsonCotacao := extrairJSON(urlDate)
	return extrairCotacaoFromJSON(jsonCotacao)
}

type estrutura struct {
	Value []struct {
		CotacaoVenda float64 `json:"cotacaoVenda"`
	} `json:"value"`
}

func extrairCotacaoFromJSON(jsonData []byte) float64 {
	var e estrutura
	json.Unmarshal(jsonData, &e)
	if len(e.Value) == 0 {
		return -1
	}
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
