package create_user_use_case

import (
	domain_user "github.com/shinYeongHyeon/onlyOurs-api/src/user/domain"
	postgresUserRepository "github.com/shinYeongHyeon/onlyOurs-api/src/user/infra/postgres"
)

// Exec: user_application_create_user_use_case.exec()
func Exec(userNewProps domain_user.UserNewProps) CreateUserUseCaseResponse {
	user := domain_user.UserNewCreate(userNewProps)

	err := postgresUserRepository.Save(user)

	if err != nil {
		return CreateUserUseCaseResponse{
			Ok: "QUERY_FAIL",
			User: user,
		}
	}

	return CreateUserUseCaseResponse{
		Ok: "SUCCESS",
		User: user,
	}
}

type CreateUserUseCaseResponse struct {
	Ok   string
	User domain_user.User
}
