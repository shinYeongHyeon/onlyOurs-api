package domain_user

import "errors"

type UserPassword struct {
	Value string
}

func UserPasswordCreate(userPasswordString string) UserPassword {
	if userPasswordString == "" {
		panic(errors.New("USER_PASSWORD_STRING WAS NOT EMPTY"))
	}

	return UserPassword {
		userPasswordString,
	}
}
