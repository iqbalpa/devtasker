package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	return string(bytes), err
}

func ComparePassword(encrypted string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(pass))
	return err == nil
}
