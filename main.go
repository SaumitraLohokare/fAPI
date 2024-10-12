package main

import (
	"fmt"
	"net/http"

	parser "github.com/SaumitraLohokare/fAPI/parser"
)

func main() {
	fileName := "default.fapi"
	var p parser.Parser
	p.ParseFile(fileName)

	fmt.Println(p.Output)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", get)
	mux.HandleFunc("POST /", post)
	mux.HandleFunc("PUT /", put)
	mux.HandleFunc("DELETE /", delete)

	server := &http.Server{
		Addr:    ":6969",
		Handler: mux,
	}

	server.ListenAndServe()
}

func get(w http.ResponseWriter, req *http.Request) {

}

func post(w http.ResponseWriter, req *http.Request) {

}

func put(w http.ResponseWriter, req *http.Request) {

}

func delete(w http.ResponseWriter, req *http.Request) {

}
