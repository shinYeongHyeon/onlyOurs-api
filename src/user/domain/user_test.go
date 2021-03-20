package domain_user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserNewCreate(t *testing.T) {
	userId := "shindevil1@naver.com"
	userName := "신영현"
	userPassword := "비밀"

	user := UserNewCreate(UserNewProps {
		userId:   userId,
		name:     userName,
		password: userPassword,
	})

	assert.Equal(t, user.userId.value, userId)
	assert.Equal(t, user.name.value, userName)
	assert.Equal(t, user.password.value, userPassword)
}

func TestUserCreate(t *testing.T) {
	id := "abcdefg"
	userId := "shindevil1@naver.com"
	userName := "신영현"
	userPassword := "비밀"

	user := UserCreate(UserProps {
		id:       id,
		userId:   userId,
		name:     userName,
		password: userPassword,
	})

	assert.Equal(t, user.id, id)
	assert.Equal(t, user.userId.value, userId)
	assert.Equal(t, user.name.value, userName)
	assert.Equal(t, user.password.value, userPassword)
}
