package postgresUserRepository

import (
	"github.com/shinYeongHyeon/onlyOurs-api/core/postgres"
	domain_user "github.com/shinYeongHyeon/onlyOurs-api/src/user/domain"
)

// Save User
func Save(user domain_user.User) error {
	_, err := postgres.PostgresDb.
		Users.
		Create().
		SetID(user.Id).
		SetUserID(user.UserId.Value).
		SetName(user.Name.Value).
		SetPassword(user.Password.Value).
		Save(postgres.PostgresCtx)

	return err
}