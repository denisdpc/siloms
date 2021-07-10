package siloms

// https://www.tesourodireto.com.br/json/br/com/b3/tesourodireto/service/api/treasurybondsinfo.json

import (
	"fmt"
	"net/http"

	"crypto/tls"
	"io/ioutil"
	"os"

	_ "github.com/gocolly/colly"
)

// Parametro contempla os parâmetros para extração de dados
type Parametro struct {
	DataInicial string
	DataAtual   string
}

type tesouro struct {
	nome string `json:"nm"`
}

func Fator(p Parametro, f func(Parametro) float64) float64 {
	return f(p)
}

func ExtrairCotacao(p Parametro) float64 {
	//c := colly.NewCollector()
	//fmt.Println(c)
	return 0.1
}

func ExtrairIGPM(p Parametro) float64 {
	return 0.2
}

func ExtrairTaxaTesouro(p Parametro) float64 {
	url := "https://www.tesourodireto.com.br/json/br/com/b3/tesourodireto/service/api/treasurybondsinfo.json"
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	response, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	htmlData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(os.Stdout, string(htmlData))

	return 0.3
}
