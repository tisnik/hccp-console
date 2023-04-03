package main

// TODO: routes as context vars?

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const (
	indexPageFilename = "templates/index.htm"
)

const (
	// ContentTypeHTML represents content type text/html used in HTTP responses
	ContentTypeHTML = "text/html"

	// ContentTypeJavaScript represents content type application/javascript used in HTTP responses
	ContentTypeJavaScript = "application/json"

	// ContentTypeCSS represents content type text/css used in HTTP responses
	ContentTypeCSS = "text/css"
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

func routeEnableHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Enabling route")

	routeID, ok := request.URL.Query()["id"]
	if !ok {
		log.Printf("Router ID not provided")
	}

	err := enableRouteWithID(routeID)
	if err != nil {
		log.Println(err)
	}

	indexPageHandler(writer, request)
}

func routeDisableHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Disabling route")

	routeID, ok := request.URL.Query()["id"]
	if !ok {
		log.Printf("Router ID not provided")
	}

	err := disableRouteWithID(routeID)
	if err != nil {
		log.Println(err)
	}

	indexPageHandler(writer, request)
}

func notFoundResponse(writer http.ResponseWriter) {
	writeResponse(writer, "Not found!")
}

func getContentType(filename string) string {
	// TODO: to map
	if strings.HasSuffix(filename, ".html") {
		return ContentTypeHTML
	} else if strings.HasSuffix(filename, ".js") {
		return ContentTypeJavaScript
	} else if strings.HasSuffix(filename, ".css") {
		return ContentTypeCSS
	}
	return ContentTypeHTML
}

func sendStaticPage(writer http.ResponseWriter, filename string) {
	// #nosec G304
	body, err := os.ReadFile(filename)
	if err == nil {
		writer.Header().Set("Server", "A Go Web Server")
		writer.Header().Set("Content-Type", getContentType(filename))
		_, err = fmt.Fprint(writer, string(body))
		if err != nil {
			log.Println("Error sending response body", err)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
		notFoundResponse(writer)
	}
}

func staticPage(filename string) func(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Serving static file %s", filename)
	return func(writer http.ResponseWriter, request *http.Request) {
		sendStaticPage(writer, filename)
	}
}

func writeResponse(writer http.ResponseWriter, message string) {
	_, err := fmt.Fprint(writer, message)
	if err != nil {
		log.Println("Error sending response", err)
	}
}

func startServer(address string) {
	// TODO: use Gorilla mux
	log.Printf("Starting HTTP server at address %s", address)

	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "application/foobar")

	// handler for templates etc
	http.HandleFunc("/", indexPageHandler)

	// handlers for all static files
	http.HandleFunc("/jquery.js", staticPage("static/jquery.js"))
	http.HandleFunc("/bootstrap.min.css", staticPage("static/bootstrap.min.css"))
	http.HandleFunc("/bootstrap.min.js", staticPage("static/bootstrap.min.js"))
	http.HandleFunc("/ccx.css", staticPage("static/ccx.css"))

	// handlers for REST API like (just like) calls
	http.HandleFunc("/route/enable", routeEnableHandler)
	http.HandleFunc("/route/disable", routeDisableHandler)

	http.ListenAndServe(address, nil)
}
