package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashNum := os.Getenv("HASH")
	num, err := strconv.Atoi(hashNum)
	if err != nil {
		return "", err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), num)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
