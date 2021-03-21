package user_controller

import (
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	user_application_create_user_use_case "github.com/shinYeongHyeon/onlyOurs-api/src/user/application"
	domain_user "github.com/shinYeongHyeon/onlyOurs-api/src/user/domain"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer core.HandleRecover(w)

	var userNewProps domain_user.UserNewProps
	core.ParseJSON(r.Body, &userNewProps)

	user_application_create_user_use_case.Exec(userNewProps)
}