package siloms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func ExtrairFatorIGPM(p Parametro) float64 {
	dtInicial := p.DataInicial.Format("01/2006")
	dtFinal := p.DataFinal.Format("01/2006")

	fmt.Println(dtInicial, dtFinal)

	url := "https://www3.bcb.gov.br/CALCIDADAO/publico/corrigirPorIndice.do?method=corrigirPorIndice"
	pesquisa := fmt.Sprintf("aba=1&selIndice=28655IGP-M&dataInicial=%s&dataFinal=%s&valorCorrecao=1", dtInicial, dtFinal)
	payload := strings.NewReader(pesquisa)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	re := regexp.MustCompile(">R\\$ &nbsp;(.*?)&nbsp;")
	f := re.FindAll(body, -1)[1]
	f = f[11 : len(f)-7]
	for i, el := range f {
		if el == 44 { // substitui , por .
			f[i] = 46
		}
	}
	fator, _ := strconv.ParseFloat(string(f), 64)

	return fator
}
