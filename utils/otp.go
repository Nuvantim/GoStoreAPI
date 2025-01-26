package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/jordan-wright/email"
	"math/big"
	"net/smtp"
	"os"
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
	mailPort := os.Getenv("MAIL_PORT")
	mailUsername := os.Getenv("MAIL_USERNAME")
	mailPassword := os.Getenv("MAIL_PASSWORD")
	mailAddress := os.Getenv("MAIL_FROM_ADDRESS")

	// set Email
	e := email.NewEmail()
	e.From = AppName + " <" + mailAddress + ">"
	e.To = []string{targetEmail} // Target email
	e.Subject = "Verify Your Account"
	e.Text = []byte("Your OTP for verification is: " + otp)

	// Combine SMTP server & port
	serverAddr := fmt.Sprintf("%s:%s", mailSMTP, mailPort)

	// Send Email
	err := e.Send(serverAddr, smtp.PlainAuth("", mailUsername, mailPassword, mailSMTP))
	if err != nil {
		return err
	}
	return nil
}
