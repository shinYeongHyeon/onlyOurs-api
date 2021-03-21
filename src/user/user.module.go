package user

import (
	"github.com/gorilla/mux"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	user_controller "github.com/shinYeongHyeon/onlyOurs-api/src/user/presentation/controller"
	"net/http"
)

// UserModule: UserModule
func UserModule() http.Handler {
	router := mux.NewRouter().PathPrefix("/user").Subrouter()

	router.NotFoundHandler = core.NotFoundHandler("User")
	router.HandleFunc("/", user_controller.CreateUser).Methods("GET")

	return router
}
