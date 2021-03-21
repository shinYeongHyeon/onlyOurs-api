package domain_user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPasswordCreate(t *testing.T) {
	userPasswordString := "pwd"
	userPassword := UserPasswordCreate(userPasswordString)
	assert.Equal(t, userPassword.Value, userPasswordString, "Create with userPassword")
}

func TestUserPasswordCreateFailByEmptyString(t *testing.T) {
	userPasswordString := ""
	assert.Panics(t, func() {
		UserPasswordCreate(userPasswordString)
	}, "panic")
}