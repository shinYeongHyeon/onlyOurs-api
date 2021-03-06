package domain_user

import "errors"

type UserId struct {
	Value string
}

func UserIdCreate(userIdString string) UserId {
	if userIdString == "" {
		panic(errors.New("USER_ID_STRING WAS NOT EMPTY"))
	}

	return UserId{
		userIdString,
	}
}
