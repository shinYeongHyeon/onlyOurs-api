package postgresUserRepository

import (
	"github.com/shinYeongHyeon/onlyOurs-api/core/postgres"
	"github.com/shinYeongHyeon/onlyOurs-api/ent"
	"github.com/shinYeongHyeon/onlyOurs-api/ent/users"
	domain_user "github.com/shinYeongHyeon/onlyOurs-api/src/user/domain"
)

func FindOne(userDomain domain_user.User) *ent.Users {
	foundUser, _ := postgres.PostgresDb.
		Users.
		Query().
		Where(users.UserID(userDomain.UserId.Value)).
		Only(postgres.PostgresCtx)

	return foundUser
}

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