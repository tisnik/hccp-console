/*
Copyright © 2023 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	indexPageFilename            = "templates/index.htm"
	errorPageFilename            = "templates/error.htm"
	haProxyConfigurationFilename = "templates/configuration.htm"
)

const (
	// ContentTypeHTML represents content type text/html used in HTTP responses
	ContentTypeHTML = "text/html"

	// ContentTypeJavaScript represents content type application/javascript used in HTTP responses
	ContentTypeJavaScript = "application/json"

	// ContentTypeCSS represents content type text/css used in HTTP responses
	ContentTypeCSS = "text/css"
)

// Messages
const (
	routerIDNotProvidedMessage    = "Router ID not provided"
	unableToRegisterMIMEExtension = "Unable to register MIME extension"
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

func errorPageHandler(writer http.ResponseWriter, request *http.Request, errorToDisplay error) {
	// re-read template (so we will be able to change template on the fly)
	// TODO: do it in init() in production code
	errorPageTemplate := template.Must(template.ParseFiles(errorPageFilename))

	log.Printf("Handling request at %s as error", request.URL.Path)

	// error string
	errorStr := errorToDisplay.Error()

	// apply template
	err := errorPageTemplate.Execute(writer, errorStr)
	if err != nil {
		log.Panic(err)
	}
}

func routeEnableHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Enabling route")

	err := request.ParseForm()
	if err != nil {
		log.Printf(routerIDNotProvidedMessage)
		errorPageHandler(writer, request, fmt.Errorf(routerIDNotProvidedMessage))
		return
	}

	routeID := request.FormValue("id")

	err = enableRouteWithID(routeID)
	if err != nil {
		log.Println(err)
		errorPageHandler(writer, request, err)
		return
	}

	indexPageHandler(writer, request)
}

func routeDisableHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Disabling route")

	err := request.ParseForm()
	if err != nil {
		log.Printf(routerIDNotProvidedMessage)
		errorPageHandler(writer, request, fmt.Errorf(routerIDNotProvidedMessage))
		return
	}

	routeID := request.FormValue("id")

	err = disableRouteWithID(routeID)
	if err != nil {
		log.Println(err)
		errorPageHandler(writer, request, err)
		return
	}

	indexPageHandler(writer, request)
}

func haProxyCheckInstallationHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Check HAProxy installation")

	exists, err := executableExist("haproxy")
	if err != nil {
		log.Println(err)
		errorPageHandler(writer, request, err)
		return
	}

	if exists {
		log.Printf("Exists")
		sendStaticPage(writer, "static/haproxy_is_installed.htm")
	} else {
		log.Printf("Does not exist")
		sendStaticPage(writer, "static/haproxy_is_not_installed.htm")
	}
}

func haProxyRunningHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Check if HAProxy is running")

	if processRunning("haproxy") {
		log.Printf("Running")
		sendStaticPage(writer, "static/haproxy_running.htm")
	} else {
		log.Printf("Not running")
		sendStaticPage(writer, "static/haproxy_not_running.htm")
	}
}

func haProxyDisplayStatusHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Display HAProxy status")
	indexPageHandler(writer, request)
}

func haProxyDisplayConfiguration(writer http.ResponseWriter, request *http.Request) {
	log.Printf("Display HAProxy configuration")
	displayConfigurationTemplate := template.Must(template.ParseFiles(haProxyConfigurationFilename))

	log.Printf("Handling request at %s as display configuration page", request.URL.Path)

	configuration, err := readHAProxyConfiguration()
	if err != nil {
		log.Panic(err)
	}

	// apply template
	err = displayConfigurationTemplate.Execute(writer, configuration)
	if err != nil {
		log.Panic(err)
	}
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

	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		log.Fatal(unableToRegisterMIMEExtension, err)
	}

	err = mime.AddExtensionType(".css", "application/foobar")
	if err != nil {
		log.Fatal(unableToRegisterMIMEExtension, err)
	}

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
	http.HandleFunc("/haproxy/check_installation", haProxyCheckInstallationHandler)
	http.HandleFunc("/haproxy/check_running", haProxyRunningHandler)
	http.HandleFunc("/haproxy/display_status", haProxyDisplayStatusHandler)
	http.HandleFunc("/haproxy/display_configuration", haProxyDisplayConfiguration)

	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Unable to start HTTP server", err)
	}
}
