package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashBycrypt(password string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash
}
