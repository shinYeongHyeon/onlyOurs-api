package core

import (
	"encoding/json"
	"net/http"
)

type coreError struct {
	message string
	status int
}

var internalError = coreError {
	"Internal Error",
	http.StatusInternalServerError,
}

func NotFoundHandler(domain string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Request was not found On " + domain + " Domain",
		})
	}
}