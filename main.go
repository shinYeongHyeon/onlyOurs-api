package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Do Enjoy!"))
	})
	http.ListenAndServe(":9999", nil)
}
