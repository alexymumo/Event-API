package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secrectKey = []byte("1234WEWWEW")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secrectKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
