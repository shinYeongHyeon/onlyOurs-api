package src

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	"github.com/shinYeongHyeon/onlyOurs-api/core/postgres"
	"github.com/shinYeongHyeon/onlyOurs-api/src/todo"
	"github.com/shinYeongHyeon/onlyOurs-api/src/user"
	"net/http"
)

// ListenAndServe starts up the server
func ListenAndServe(address string, r *mux.Router) error {
	postgres.Connect()

	http.Handle("/", core.LoggingMiddleware(r))
	http.Handle("/todo/", core.LoggingMiddleware(todo.TodoModule()))
	http.Handle("/user/", core.LoggingMiddleware(user.UserModule()))

	return http.ListenAndServe(
		address,
		nil,
	)
}
