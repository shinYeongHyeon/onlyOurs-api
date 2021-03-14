package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	_ "github.com/shinYeongHyeon/onlyOurs-api/docs"
	"github.com/shinYeongHyeon/onlyOurs-api/ent"
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
	url := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		"sampleuser",
		"samplesecret",
		"onlyours-postgres",
		"5432",
		"sampledb")
	client, err := ent.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

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
