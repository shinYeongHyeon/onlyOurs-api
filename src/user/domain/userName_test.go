package domain_user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserNameCreate(t *testing.T) {
	userNameString := "신영현"
	userName := UserNameCreate(userNameString)
	assert.Equal(t, userName.Value, userNameString, "Create with userName")
}

func TestUserNameCreateFailByEmptyString(t *testing.T) {
	userNameString := ""
	assert.Panics(t, func() {
		UserNameCreate(userNameString)
	}, "panic")
}