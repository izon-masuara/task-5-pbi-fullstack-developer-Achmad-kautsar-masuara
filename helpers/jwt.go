package helpers

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/izon-masuara/app"
)

func GenerateToken(users app.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID": users.ID,
	})

	resToken, err := token.SignedString([]byte(os.Getenv("JWT")))
	if err != nil {
		return "", err
	}
	return resToken, nil
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT")), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
