package siloms

import (
	"crypto/tls"
	"encoding/json"
	"strings"

	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "regexp"
	_ "strings"
)

func ExtrairTesouro(p Parametro) float64 {
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

	TrsrBondMkt(jsonData)

	return 0.7
}

func TrsrBondMkt(jsonData []byte) {
	type Master struct {
		Response struct {
			TrsrBdTradgList []struct {
				TrsrBd struct {
					Nm              string  `json:"nm"`
					AnulInvstmtRate float32 `json:"anulInvstmtRate"`
				} `json:"TrsrBd"`
			} `json:"TrsrBdTradgList"`
		} `json:"response"`
	}
	var m Master
	json.Unmarshal(jsonData, &m)

	var prefixado string
	var taxaMin float32

	for _, v := range m.Response.TrsrBdTradgList {
		titulo, taxa := v.TrsrBd.Nm, v.TrsrBd.AnulInvstmtRate
		if strings.HasPrefix(titulo, "Tesouro Prefixado 20") && taxa > 0 {
			prefixado = titulo
			taxaMin = taxa
		}

	}
	fmt.Println(prefixado, taxaMin)
	//fmt.Println(m.Response.TrsrBdTradgList[0])
}
