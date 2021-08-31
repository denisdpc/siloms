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
)

// Servidor gerenciamento de requisições
func Servidor() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../siloms/web/upload.html"))
	tmpl.Execute(w, nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // maximum upload of 10 MB files

	fileAtual := criarArquivo("fileAtual", r)
	fileRef := criarArquivo("fileRef", r)

	mapaPNtoReqsRef := siloms.MapaPNtoRequisicoes(fileRef)
	mapaPNtoReqsAtual := siloms.MapaPNtoRequisicoes(fileAtual)

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

	// type Todo struct {
	// 	Title string
	// 	Done  bool
	// }

	// type TodoPageData struct {
	// 	PageTitle string
	// 	Todos     []Todo
	// }

	// data := TodoPageData{
	// 	PageTitle: "My TODO list",
	// 	Todos: []Todo{
	// 		{Title: "Task 1", Done: false},
	// 		{Title: "Task 2", Done: true},
	// 		{Title: "Task 3", Done: true},
	// 	},
	// }

	tmpl := template.Must(template.ParseFiles("../siloms/web/correspondencia.html"))
	tmpl.Execute(w, data)
	//tmpl.Execute(w, nil)

	// TODO: direcionar para página de correspondencia
	// com as correspondentes referências para escolha da mais compatível

	//http.Redirect(w, r, "/", http.StatusFound)

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
