package web

// https://tutorialedge.net/golang/go-file-upload-tutorial/
// https://golang.org/doc/
// https://stackoverflow.com/questions/40684307/how-can-i-receive-an-uploaded-file-using-a-golang-net-http-server

import (
	"fmt"
	"github.com/denisdpc/siloms"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

// Servidor gerenciamento de requisições
func Servidor() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/planilhar", planilhar)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../siloms/web/upload.html"))
	tmpl.Execute(w, nil)
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// saida: reqRef e []Requisicao atuais
func planilhar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for partNumber, numReqRef := range r.Form {
		var reqRef siloms.Requisicao
		var reqs []siloms.Requisicao

		for _, req := range mapaPNtoReqsRef[partNumber] {
			if req.Numero == numReqRef[0] {
				reqs = mapaPNtoReqsAtual[partNumber]
				reqRef = req

				p := &siloms.Parametro{
					DataInicial: req.DataPlano,
					DataFinal:   time.Now().AddDate(0, -1, 0)}

				pIGPM := siloms.ExtrairFatorIGPM(p)
				p.DataFinal = time.Now()
				pCot, _, _ := siloms.ExtrairFatorCotacao(p)
				pTDPre, titulo := siloms.ExtrairTaxaTesouro()
				caminho := "../csv/"
				R := 0.05

				siloms.Planilhar(caminho, time.Now(), reqRef, reqs, pTDPre, pCot, pIGPM, R)
				fmt.Println(titulo)

				break
			}
		}
	}
	// TODO: redirecionar para página que permite realizar o download dos arquivos
}

var mapaPNtoReqsRef map[string][]siloms.Requisicao
var mapaPNtoReqsAtual map[string][]siloms.Requisicao

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // maximum upload of 10 MB files

	fileAtual := criarArquivo("fileAtual", r)
	fileRef := criarArquivo("fileRef", r)

	mapaPNtoReqsRef = siloms.MapaPNtoRequisicoes(fileRef)
	mapaPNtoReqsAtual = siloms.MapaPNtoRequisicoes(fileAtual)
	//mapaPNtoReqsRef := siloms.MapaPNtoRequisicoes(fileRef)
	//mapaPNtoReqsAtual := siloms.MapaPNtoRequisicoes(fileAtual)

	var reqsPnRefToAtual []siloms.RequisicaoPnRefToAtual
	for partNumber, reqs := range mapaPNtoReqsRef {
		reqRefAtual := siloms.RequisicaoPnRefToAtual{
			PartNumber: partNumber,
			ReqsRef:    reqs,
			ReqsAtual:  mapaPNtoReqsAtual[partNumber]}
		reqsPnRefToAtual = append(reqsPnRefToAtual, reqRefAtual)
	}

	type Data struct {
		Lista []siloms.RequisicaoPnRefToAtual
	}
	data := Data{Lista: reqsPnRefToAtual}

	tmpl := template.Must(template.ParseFiles("../siloms/web/correspondencia.html"))
	tmpl.Execute(w, data)
}

func criarArquivo(arqNome string, r *http.Request) string {
	file, _, err := r.FormFile(arqNome)
	if err != nil {
		fmt.Println("erro ao recuperar o arquivo " + arqNome)
		panic(err)
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("../csv/", arqNome+"-*.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	return tempFile.Name()
}
