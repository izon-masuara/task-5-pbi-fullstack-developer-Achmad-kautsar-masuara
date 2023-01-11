package helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/izon-masuara/app"
)

func GenerateToken(users app.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID": users.ID,
	})

	resToken, err := token.SignedString([]byte("AccessToekn"))
	if err != nil {
		return "", err
	}
	return resToken, nil
}
