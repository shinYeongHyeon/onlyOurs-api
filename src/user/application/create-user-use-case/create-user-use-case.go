package create_user_use_case

import (
	"github.com/shinYeongHyeon/onlyOurs-api/core/postgres"
	domain_user "github.com/shinYeongHyeon/onlyOurs-api/src/user/domain"
)

// Exec: user_application_create_user_use_case.exec()
func Exec(userNewProps domain_user.UserNewProps) CreateUserUseCaseResponse {
	user := domain_user.UserNewCreate(userNewProps)

	_, err := postgres.PostgresDb.
		Users.
		Create().
		SetID(user.Id).
		SetUserID(user.UserId.Value).
		SetName(user.Name.Value).
		SetPassword(user.Password.Value).
		Save(postgres.PostgresCtx)

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
