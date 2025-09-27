package helper

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var secretKeyJwt = os.Getenv("JWT_SECRET_KEY")

func CreateToken(userId, roleId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"roleId": roleId,
	})

	tokenString, err := token.SignedString([]byte(secretKeyJwt))

	return tokenString, err
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKeyJwt, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
