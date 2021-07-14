package siloms

import (
	"crypto/tls"
	"encoding/json"
	"strings"

	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ExtrairTaxaTesouro obtém a taxa de retorno do tesouro pré-fixado com vencimento mais próximo
func ExtrairTaxaTesouro() float64 {
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

	jsonData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return trsrBondMkt(jsonData)
}

func trsrBondMkt(jsonData []byte) float64 {
	type estrutura struct {
		Response struct {
			TrsrBdTradgList []struct {
				TrsrBd struct {
					Nm              string  `json:"nm"`
					AnulInvstmtRate float64 `json:"anulInvstmtRate"`
				} `json:"TrsrBd"`
			} `json:"TrsrBdTradgList"`
		} `json:"response"`
	}
	var e estrutura
	json.Unmarshal(jsonData, &e)

	var prefixado string
	var taxaMin float64

	for _, v := range e.Response.TrsrBdTradgList {
		titulo, taxa := v.TrsrBd.Nm, v.TrsrBd.AnulInvstmtRate
		if strings.HasPrefix(titulo, "Tesouro Prefixado 20") && taxa > 0 {
			prefixado = titulo
			taxaMin = taxa
		}

	}
	fmt.Println(prefixado, taxaMin)
	return taxaMin
}
