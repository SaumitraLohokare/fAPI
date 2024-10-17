package main

import (
	"fmt"
	"net/http"

	"github.com/SaumitraLohokare/fAPI/parser"
)

type FAPI_Handler struct {
	Url       string
	Responses map[string]parser.Request
}

// ServeHTTP implements http.Handler.
func (h FAPI_Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response, ok := h.Responses[r.Method]
	if !ok {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintf("%s method not implementd for %s.", r.Method, h.Url)))
	}

	w.WriteHeader(response.Status)
	w.Header().Add("Content-Type", response.ContentType)
	w.Write([]byte(response.Response))
}

func GenerateHandlers(ctx *parser.Context) {
	routes := ctx.Routes

	for _, route := range routes {
		methodsMap := make(map[string]parser.Request)

		if isMethodValid(route.Get) {
			methodsMap["GET"] = route.Get
		}
		if isMethodValid(route.Post) {
			methodsMap["POST"] = route.Post
		}
		// TODO: as more methods get supported add appropriate statements here

		handler := FAPI_Handler{
			Url:       route.Url,
			Responses: methodsMap,
		}

		http.Handle(route.Url, handler)
	}
}

// If this method was not specified in the JSON, it's status defaults to 0
func isMethodValid(req parser.Request) bool {
	return req.Status != 0
}
