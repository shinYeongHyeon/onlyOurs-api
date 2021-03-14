package domain_user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserEmailCreate(t *testing.T) {
	userEmailString := "신영현"
	userEmail := UserEmailCreate(userEmailString)
	assert.Equal(t, userEmail.value, userEmailString, "Create with userEmail")
}

func TestUserEmailCreateFailByEmptyString(t *testing.T) {
	userEmailString := ""
	assert.Panics(t, func() {
		UserEmailCreate(userEmailString)
	}, "panic")
}