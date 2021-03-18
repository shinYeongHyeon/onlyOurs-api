package domain_user

type User struct {
	id string
	userId UserId
	name UserName
	password UserPassword
}

type UserProps struct {
	id string
	userId string
	name string
	password string
}

type UserNewProps struct {
	userId string
	name string
	password string
}

func UserCreate(userProps UserProps) User {
	return User {
		id: userProps.id,
		userId: UserIdCreate(userProps.userId),
		name: UserNameCreate(userProps.name),
		password: UserPasswordCreate(userProps.password),
	}
}

func UserNewCreate(userNewProps UserNewProps) User {
	return UserCreate(UserProps {
		id:       "todo uuid",
		userId:   userNewProps.userId,
		name:     userNewProps.name,
		password: userNewProps.password,
	})
}