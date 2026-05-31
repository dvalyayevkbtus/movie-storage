package main

import (
	"io"
	"net/http"
)

func MethodNotAllowed(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusMethodNotAllowed)
	io.WriteString(rw, "Method not allowed!")
}

func InternalServerError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	io.WriteString(rw, "Internal server error!")
}

func SuccessString(rw http.ResponseWriter, body string) {
	io.WriteString(rw, body)
}
