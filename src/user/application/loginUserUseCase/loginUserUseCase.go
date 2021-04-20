package loginUserUseCase

import "github.com/labstack/gommon/log"

// Exec: login
func Exec(loginUserUseCaseRequestBody LoginUserUseCaseRequestBody) error {
	log.Print(loginUserUseCaseRequestBody)

	return nil
}
