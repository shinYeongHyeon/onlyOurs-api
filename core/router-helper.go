package core

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"reflect"
)

// Must : Check Internal Error
func Must(err error) {
	if err != nil {
		log.Error("Internal error : ", err)
		panic(internalError)
	}
}

// ParseJSON
func ParseJSON(r io.Reader, v interface{}) {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		log.Error("Error In parsing json body : ", err)
		panic(malformedInputError)
	}
}

// WriteJSON : Response Api Call
func WriteJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	Must(json.NewEncoder(w).Encode(v))
}

// HandleRecover
func HandleRecover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		status := 500
		message := ""

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		switch r.(type) {
		case Error:
			errorObject := reflect.ValueOf(r).Interface().(Error)
			status = errorObject.Status
			message = errorObject.Message
		default:
		}

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": message,
		})
	}
}