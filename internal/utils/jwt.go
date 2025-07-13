package utils

import (
	"devtasker/internal/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user model.User) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":     user.Name,
			"username": user.Username,
			"exp":      time.Now().Add(time.Duration(0.5 * float64(time.Hour))).Unix(),
		})
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	key := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	return nil, false
}
