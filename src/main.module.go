package src

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	"github.com/shinYeongHyeon/onlyOurs-api/core/postgres"
	"github.com/shinYeongHyeon/onlyOurs-api/src/todo"
	"net/http"
)

type test2 struct {
	A int
	B string
}

// ListenAndServe starts up the server
func ListenAndServe(address string, r *mux.Router) error {
	postgres.Connect()

	http.Handle("/", core.LoggingMiddleware(r))
	http.Handle("/todo/", core.LoggingMiddleware(todo.TodoModule()))

	return http.ListenAndServe(
		address,
		nil,
	)
}
