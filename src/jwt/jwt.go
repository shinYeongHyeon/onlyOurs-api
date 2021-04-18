package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SCRET_KEY")))

	if err != nil {
		return "", err
	}

	return token, nil
}
