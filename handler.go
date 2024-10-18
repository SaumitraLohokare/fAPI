package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/SaumitraLohokare/fAPI/parser"
)

type FAPI_Handler struct {
	Url       string
	Responses map[string]parser.Request
	UrlRegExp string
}

// ServeHTTP implements http.Handler.
func (h FAPI_Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	match, err := regexp.Match(h.UrlRegExp, []byte(r.URL.String()))
	if err != nil {
		log.Default().Printf("| %s : %s -> %d", r.URL.String(), r.Method, 500)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("[ERROR]: %s.", err)))
		return
	}

	if !match {
		log.Default().Printf("| %s : %s -> %d", r.URL.String(), r.Method, 404)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%s route is not implemented.", r.URL)))
		return
	}

	response, ok := h.Responses[r.Method]
	if !ok {
		log.Default().Printf("| %s : %s -> %d", r.URL, r.Method, 500)
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintf("%s method not implementd for %s.", r.Method, h.Url)))
		return
	}

	log.Default().Printf("| %s : %s -> %d", r.URL, r.Method, response.Status)
	w.Header().Set("Content-Type", response.ContentType)
	w.WriteHeader(response.Status)
	w.Write([]byte(response.Response))
}

// Generates handlers for all routes mentioned in the JSON
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
		if isMethodValid(route.Put) {
			methodsMap["PUT"] = route.Put
		}
		if isMethodValid(route.Delete) {
			methodsMap["DELETE"] = route.Delete
		}
		// TODO: as more methods get supported add appropriate statements here

		urlRegExp := MakeRegExpFromUrl(route.Url)
		handler := FAPI_Handler{
			Url:       route.Url,
			Responses: methodsMap,
			UrlRegExp: urlRegExp,
		}

		http.Handle(route.Url, handler)
	}
}

// If this method was not specified in the JSON, it's status defaults to 0
func isMethodValid(req parser.Request) bool {
	return req.Status != 0
}

// Makes a Regular Expression from a given URL
// Supports '*' Wildcard
// Eg. /users/*/ matches /users/123/
func MakeRegExpFromUrl(url string) string {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	pattern := "^" + url + "$"
	return pattern
}
