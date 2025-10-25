package auth

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// hash password and return
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

func ValidatePasswordWithHash(password string, hash string) bool {
	byte_password := []byte(password)

	err := bcrypt.CompareHashAndPassword([]byte(hash), byte_password)

	return err == nil
}

func GenerateToken(length int) (string, error) {
	bytes := make([]byte, length)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), err
}

// updating token: cookie in main code, db in db code
