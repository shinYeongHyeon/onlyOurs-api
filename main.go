package main

import (
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	_ "github.com/shinYeongHyeon/onlyOurs-api/docs"
	"github.com/shinYeongHyeon/onlyOurs-api/src"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title onlyOurs API
// @version 0.0.1
// @description This is onlyOurs API
// @host localhost:9999
// @BasePath /
func main() {
	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	err := src.ListenAndServe(":9999", router)
	if err != nil {
		log.Fatal("Error On Server")
	}
}
