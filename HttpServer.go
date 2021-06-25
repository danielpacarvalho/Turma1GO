package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"text/template"
)

var path string = "/home/danielcarvalho/Documents/BoxTurtle/"
var templates = template.Must(template.ParseFiles(path + "upload.html"))

func display(w http.ResponseWriter, pagina string, dados interface{}) {
	caminho := pagina + ".html"
	templates.ExecuteTemplate(w, caminho, dados)
}

func main() {

	log.Println("Iniciando sistema BoxTurtle...")
	log.Println("Iniciado com sucesso, acesse em http://localhost:3000/upload ")
	http.HandleFunc("/upload", receberArquivo)
	http.ListenAndServe(":3000", nil)
}

func receberArquivo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "upload", nil)
	case "POST":
		uploadFile(w, r)
	}

}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, fileHandler, err := r.FormFile("arquivoEntrada")
	if err != nil {
		fmt.Println("Erro ao receber arquivo")
		fmt.Println(err)
	}
	defer file.Close()
	if gravarBanco(file, fileHandler) {
		fmt.Fprintf(w, "Arquivo gravado no banco com sucesso\n")
	} else {
		fmt.Fprintf(w, "Erro ao gravar o arquivo\n")
	}

}

type Pessoa struct {
	Nome   string
	Idade  int
	Comida string
}

func gravarBanco(file multipart.File, fileHandler *multipart.FileHeader) bool {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	pessoa := Pessoa{
		Nome:   "Katrine",
		Idade:  7,
		Comida: "Feijão",
	}
	av, err := dynamodbattribute.MarshalMap(pessoa)
	if err != nil {
		return false
	}

	tableName := "TesteGo"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	svc := dynamodb.New(sess)
	_, err = svc.PutItem(input)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Adicionado o Item " + pessoa.Nome + " que gosta de " + pessoa.Comida + " na table " + tableName)
	}
	return err == nil
}
