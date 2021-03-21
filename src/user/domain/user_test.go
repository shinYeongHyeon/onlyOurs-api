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
		UserId:   userId,
		Name:     userName,
		Password: userPassword,
	})

	assert.Equal(t, user.UserId.Value, userId)
	assert.Equal(t, user.Name.Value, userName)
	assert.Equal(t, user.Password.Value, userPassword)
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

	assert.Equal(t, user.Id, id)
	assert.Equal(t, user.UserId.Value, userId)
	assert.Equal(t, user.Name.Value, userName)
	assert.Equal(t, user.Password.Value, userPassword)
}
