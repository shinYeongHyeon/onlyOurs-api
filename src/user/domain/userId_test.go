package domain_user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserIdCreate(t *testing.T) {
	userIdString := "신영현"
	userId := UserIdCreate(userIdString)
	assert.Equal(t, userId.value, userIdString, "Create with userId")
}

func TestUserIdCreateFailByEmptyString(t *testing.T) {
	userIdString := ""
	assert.Panics(t, func() {
		UserIdCreate(userIdString)
	}, "panic")
}