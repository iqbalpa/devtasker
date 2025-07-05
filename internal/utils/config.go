package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env file")
		return
	}
	fmt.Println("loaded .env successfully")
}
