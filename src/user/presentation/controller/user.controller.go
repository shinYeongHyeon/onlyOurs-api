package user_controller

import (
	"github.com/labstack/gommon/log"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	"github.com/shinYeongHyeon/onlyOurs-api/src/user/application/create-user-use-case"
	domain_user "github.com/shinYeongHyeon/onlyOurs-api/src/user/domain"
	"net/http"
)

type createUserResponse struct {
	Id     string
	UserId string
	Name   string
}

type LoginUserResponse struct {
	token string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer core.HandleRecover(w)

	var userNewProps domain_user.UserNewProps
	core.ParseJSON(r.Body, &userNewProps)

	response := create_user_use_case.Exec(userNewProps)
	core.WriteJSON(w, createUserResponse {
		Id:     response.User.Id,
		UserId: response.User.UserId.Value,
		Name:   response.User.Name.Value,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer core.HandleRecover(w)

	log.Print("lgon")

	// Todo: Params

	core.WriteJSON(w, LoginUserResponse {
		token: "token",
	})
}