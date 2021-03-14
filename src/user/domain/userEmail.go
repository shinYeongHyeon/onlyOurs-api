package domain_user

import "errors"

type UserEmail struct {
	value string
}

func UserEmailCreate(userEmailString string) UserEmail {
	if userEmailString == "" {
		panic(errors.New("USER_EMAIL_STRING WAS NOT EMPTY"))
	}

	return UserEmail {
		userEmailString,
	}
}
