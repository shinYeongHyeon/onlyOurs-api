package user

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	user_application_create_user_use_case "github.com/shinYeongHyeon/onlyOurs-api/src/user/application"
	"net/http"
)

// UserModule: UserModule
func UserModule() http.Handler {
	router := mux.NewRouter().PathPrefix("/user").Subrouter()

	router.NotFoundHandler = core.NotFoundHandler("User")
	router.HandleFunc("/", user_application_create_user_use_case.Exec).Methods("GET")

	return router
}
