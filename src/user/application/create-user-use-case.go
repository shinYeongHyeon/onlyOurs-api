package user_application_create_user_use_case

import domain_user "github.com/shinYeongHyeon/onlyOurs-api/src/user/domain"

// Exec: user_application_create_user_use_case.exec()
func Exec(userNewProps domain_user.UserNewProps) CreateUserUseCaseResponse {
	user := domain_user.UserNewCreate(userNewProps)

	return CreateUserUseCaseResponse {
		ok: "SUCCESS",
		user: user,
	}
}

type CreateUserUseCaseResponse struct {
	ok   string
	user domain_user.User
}
