package utils

import (
	"crypto/rand"
	"gopkg.in/gomail.v2"
	"math/big"
	"os"
	"strconv"
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

func SendOTP(targetEmail, otp string) error {
	// environment variables
	AppName := os.Getenv("APP_NAME")
	mailSMTP := os.Getenv("MAIL_MAILER")
	mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	mailUsername := os.Getenv("MAIL_USERNAME")
	mailPassword := os.Getenv("MAIL_PASSWORD")
	mailAddress := os.Getenv("MAIL_FROM_ADDRESS")

	// Create new message
	m := gomail.NewMessage()
	m.SetHeader("From", AppName+" <"+mailAddress+">")
	m.SetHeader("To", targetEmail)
	m.SetHeader("Subject", "Verify Your Account")
	m.SetBody("text/plain", "Your OTP for verification is: "+otp)

	// Create dialer
	d := gomail.NewDialer(mailSMTP, mailPort, mailUsername, mailPassword)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
