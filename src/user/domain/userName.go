package domain_user

import "errors"

type UserName struct {
	Value string
}

func UserNameCreate(userNameString string) UserName {
	if userNameString == "" {
		panic(errors.New("USER_NAME_STRING WAS NOT EMPTY"))
	}

	return UserName {
		userNameString,
	}
}
