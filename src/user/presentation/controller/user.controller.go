package user_controller

import (
	"github.com/labstack/gommon/log"
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	"github.com/shinYeongHyeon/onlyOurs-api/src/user/application/CreateUserUseCase"
	loginUserUseCase "github.com/shinYeongHyeon/onlyOurs-api/src/user/application/loginUserUseCase"
	domain_user "github.com/shinYeongHyeon/onlyOurs-api/src/user/domain"
	"net/http"
)

type createUserResponse struct {
	Code   string
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

	response := CreateUserUseCase.Exec(userNewProps)
	log.Print(response)
	core.WriteJSON(w, createUserResponse {
		Code:   response.Ok,
		Id:     response.User.Id,
		UserId: response.User.UserId.Value,
		Name:   response.User.Name.Value,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer core.HandleRecover(w)

	var loginUserUseCaseRequestBody loginUserUseCase.LoginUserUseCaseRequestBody
	core.ParseJSON(r.Body, &loginUserUseCaseRequestBody)

	response := loginUserUseCase.Exec(loginUserUseCaseRequestBody)
	log.Print(response)
	// Todo: Params

	core.WriteJSON(w, LoginUserResponse {
		token: "token",
	})
}