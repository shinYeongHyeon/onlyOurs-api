package loginUserUseCase

type LoginUserUseCaseRequestBody struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}