package lib

import (
	"golang.org/x/crypto/bcrypt"
)

// hash the given password
func HashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// check if the given password matches the hashed password
func CheckPasswordHash(p, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}
