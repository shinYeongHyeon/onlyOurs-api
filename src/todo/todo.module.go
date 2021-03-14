package todo

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	"net/http"
)

type test struct {
	A int
	B string
}

// TodoModule: TodoModules
func TodoModule() http.Handler {
	router := mux.NewRouter().PathPrefix("/todo").Subrouter()

	router.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		core.WriteJSON(w, test {
			6,
			"compose_test",
		})
	}).Methods("GET")

	return router
}
