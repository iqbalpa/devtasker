package utils

import (
	"devtasker/internal/model"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user model.User) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":     user.Name,
			"username": user.Username,
		})
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}
