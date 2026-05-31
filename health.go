package main

import "net/http"

func HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		MethodNotAllowed(rw)
		return
	}

	SuccessString(rw, "UP")
}
