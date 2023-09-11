package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"text/template"
)

type DataTemplate struct {
	FileBinary []byte
}

func receiveFile(res http.ResponseWriter, req *http.Request) (string, []byte, error) {
	var buf bytes.Buffer
	file, header, err := req.FormFile("fileToRead")

	if err != nil {
		log.Println("aqui?")
		return "", nil, err
	}

	defer file.Close()

	io.Copy(&buf, file)
	contents := buf.Bytes()
	//fmt.Println(contents)

	return header.Filename, contents, err
}
func formPage(res http.ResponseWriter, req *http.Request) {
	templateName := "main.html"
	tmplGet := template.Must(template.ParseFiles(templateName))
	tmplGet.Execute(res, nil)
}

func createTemplate(res http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodPost:
		log.Println("aquiii")
		http.Redirect(res, req, "/file", http.StatusFound)

	default:
		formPage(res, req)

	}

}

func fileChoices(res http.ResponseWriter, req *http.Request) {
	log.Println("olá")
	templateName := "filePage.html"
	fileName, bFile, err := receiveFile(res, req)
	if err != nil {
		log.Println(err)
	}
	log.Println(fileName)

	dataTemplate := DataTemplate{
		FileBinary: bFile,
	}
	templWithFile := template.Must(template.ParseFiles(templateName))

	templWithFile.Execute(res, dataTemplate)

}
func main() {
	http.HandleFunc("/", createTemplate)
	http.HandleFunc("/file", fileChoices)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
