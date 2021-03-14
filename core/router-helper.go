package core

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"net/http"
)

// Must : Check Internal Error
func Must(err error) {
	if err != nil {
		log.Error("Internal error : ", err)
		panic(internalError)
	}
}

// WriteJSON : Response Api Call
func WriteJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	Must(json.NewEncoder(w).Encode(v))
}
