package main

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	_ "github.com/shinYeongHyeon/onlyOurs-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type test struct {
	A int
	B string
}

// @title onlyOurs API
// @version 0.0.1
// @description This is onlyOurs API
// @host localhost:9999
// @BasePath /
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		core.WriteJSON(w, test {
			6,
			"compose_test",
		})
	})

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":9999", router)
}
