package core

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string
	Status int
}

var internalError = Error {
	"Internal Error",
	http.StatusInternalServerError,
}

var malformedInputError = Error {
	"Malformed Input Error",
	http.StatusBadRequest,
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