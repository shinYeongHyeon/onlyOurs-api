package main

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	"net/http"
)

type test struct {
	A int
	B string
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		core.WriteJSON(w, test {
			3,
			"test",
		})
	})

	http.ListenAndServe(":9999", router)
}
