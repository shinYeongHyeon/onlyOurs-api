package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Do Enjoy!@#"))
	})

	http.ListenAndServe(":9999", router)
}
