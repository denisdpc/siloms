package siloms

// https://www.tesourodireto.com.br/json/br/com/b3/tesourodireto/service/api/treasurybondsinfo.json
// https://jsonformatter.org/
// https://tutorialedge.net/golang/parsing-json-with-golang/
// https://www.golangprograms.com/how-to-unmarshal-nested-json-structure.html
// https://blog.serverbooter.com/post/parsing-nested-json-in-go/

import (
	"encoding/json"
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

type tesouro0 struct {
	status string `json:"responseStatusText"`
}

type tesouro1 struct {
	resposta tesouro2 `json:"response"`
}

type tesouro2 struct {
	listaTitulos []tesouro3 `json:"TrsrBdTradgList"`
}

type tesouro3 struct {
	titulo      string `json:"nm"`
	valorMinimo string `json:"minInvstmtAmt"`
	rendimento  string `json:"anulInvstmtRate"`
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

	var tes tesouro0

	json.Unmarshal(htmlData, &tes)
	fmt.Println(tes)

	fmt.Println(htmlData)

	//fmt.Println(os.Stdout, string(htmlData))
	// response --> TrsrBdTradgList --> []TrsrBd --> nm, minInvstmtAmt>0, anulInvstmtRate
	return 0.3
}

func ExtrairTeste() {

	htmlData := []byte(`{
  "responseStatus": 200,
  "responseStatusText": "success",
  "statusInfo": "OK",
  "response": {
    "BdTxTp": {
	  "cd": 1,
	  "md": 3
    },
    "TrsrBondMkt": {
      "opngDtTm": "2021-07-12T09:15:00",
      "clsgDtTm": "2021-07-13T05:00:00",
      "qtnDtTm": "2021-07-12T09:25:00.687",
      "stsCd": 3,
      "sts": "Em manutenção"
	},
	"TrsrBdTradgList": [
		{
			"TrsrBd": {
				"cd": 159,
          		"nm": "Tesouro Selic 2023"
			},
			"TrsrBd": {
				"cd": 260,
          		"nm": "Tesouro Selic 2025"
			},
			"TrsrBd": {
				"cd": 360,
          		"nm": "Tesouro Selic 2035"
			}
		}
	]
  }
}`)

	type N31 struct {
		StsCd int    `json:"stsCd"`
		Sts   string `json:"sts"`
	}

	type N42 struct {
		Cd int    `json:"cd"`
		Nm string `json:"nm"`
	}

	type N32 struct {
		TrsrBds N42 `json:"TrsrBd"`
	}

	type N21 struct {
		TrsrBondMkt     N31   `json:"TrsrBondMkt"`
		TrsrBdTradgList []N32 `json:"TrsrBdTradgList"`
	}

	type N1 struct {
		ResponseStatus     int    `json:"responseStatus"`
		ResponseStatusText string `json:"responseStatusText"`
		StatusInfo         string `json:"statusInfo"`
		Response           N21    `json:"response"`
	}

	/*
			type Foo struct {
				More String `json:"more"`
				Foo  struct {
					Bar string `json:"bar"`
					Baz string `json:"baz"`
				} `json:"foo"`
		    //  FooBar  string `json:"foo.bar"`
			}
	*/

	// https://medium.com/@xcoulon/nested-structs-in-golang-2c750403a007

	/*
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

			if etype Master struct {
			Response struct {
				TrsrBondMkt struct {
					StsCd int    `json:"stsCd"`
					Sts   string `json:"sts"`
				} `json:"TrsrBondMkt"`
			} `json:"response"`
		}rr != nil {
				fmt.Println(err)
				os.Exit(1)
			}

	*/

	/*
		type TrsrBd struct {
			Cd int    `json:"Cd"`
			Nm string `json:"nm"`
		} `json:"TrsrBd"`
	*/

	// https://medium.com/@alain.drolet.0/how-to-unmarshal-an-array-of-json-objects-of-different-types-into-a-go-struct-10eab5f9a3a2
	// https://dev.to/m7shapan/golang-how-to-unmarshal-a-subset-of-nested-json-data-d84
	// https://mariadesouza.com/2017/09/07/custom-unmarshal-json-in-golang/
	// https://mholt.github.io/json-to-go/
	// https://play.golang.org/p/e15aNv2e4gs
	type Master struct {
		Response struct {
			TrsrBondMkt struct {
				StsCd int    `json:"stsCd"`
				Sts   string `json:"sts"`
			} `json:"TrsrBondMkt"`
			TrsrBdTradgList []struct {
				TrsrBd struct {
					Cd int    `json:"cd"`
					Nm string `json:"nm"`
				} `json:"TrsrBd"`
			} `json:"TrsrBdTradgList"`
		} `json:"response"`
	}

	var jsonTexto = []byte(htmlData)

	//var t1 N1
	var t1 Master
	err := json.Unmarshal(jsonTexto, &t1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t1.Response.TrsrBdTradgList[0])

}
