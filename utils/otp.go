package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP() string {
	charset := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	code := make([]byte, 7)

	for i := range code {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		code[i] = charset[n.Int64()]
	}

	return string(code)
}
