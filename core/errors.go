package core

import "net/http"

type coreError struct {
	message string
	status int
}

var internalError = coreError {
	"Internal Error",
	http.StatusInternalServerError,
}