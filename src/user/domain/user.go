package domain_user

import "github.com/google/uuid"

type User struct {
	Id       string
	UserId   UserId
	Name     UserName
	Password UserPassword
}

type UserProps struct {
	id       string
	userId   string
	name     string
	password string
}

type UserNewProps struct {
	UserId   string `json:"userId"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func UserCreate(userProps UserProps) User {
	return User {
		Id:       userProps.id,
		UserId:   UserIdCreate(userProps.userId),
		Name:     UserNameCreate(userProps.name),
		Password: UserPasswordCreate(userProps.password),
	}
}

func UserNewCreate(userNewProps UserNewProps) User {
	return UserCreate(UserProps {
		id:       uuid.NewString(),
		userId:   userNewProps.UserId,
		name:     userNewProps.Name,
		password: userNewProps.Password,
	})
}