package web

// https://tutorialedge.net/golang/go-file-upload-tutorial/
// https://golang.org/doc/
// https://stackoverflow.com/questions/40684307/how-can-i-receive-an-uploaded-file-using-a-golang-net-http-server

import (
	"fmt"
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

	criarArquivo("fileAtual", r)
	criarArquivo("fileRef", r)

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func criarArquivo(arqNome string, r *http.Request) {
	file, _, err := r.FormFile(arqNome)
	if err != nil {
		fmt.Println("erro ao recuperar o arquivo " + arqNome)
		fmt.Println(err)
		return
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
}
