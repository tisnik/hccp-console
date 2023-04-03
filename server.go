package main

import (
	//"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	indexPageFilename = "templates/index.htm"
)

func indexPageHandler(writer http.ResponseWriter, request *http.Request) {
	// re-read template (so we will be able to change template on the fly)
	// TODO: do it in init() in production code
	indexPageTemplate := template.Must(template.ParseFiles(indexPageFilename))

	log.Printf("Handling request at %s", request.URL.Path)

	// apply template
	err := indexPageTemplate.Execute(writer, routes)
	if err != nil {
		log.Panic(err)
	}
}

func startServer(address string) {

	log.Printf("Starting HTTP server at address %s", address)

	http.HandleFunc("/", indexPageHandler)
	http.Handle("/static/", http.FileServer(http.Dir("static/")))

	http.ListenAndServe(address, nil)
}
