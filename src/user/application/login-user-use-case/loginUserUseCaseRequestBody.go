package login_user_use_case

type LoginUserUseCaseRequestBody struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}