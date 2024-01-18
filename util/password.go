package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashed), nil
}

func CheckPassword(pass, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
}
