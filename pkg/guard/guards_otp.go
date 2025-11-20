package guard

import (
	"crypto/rand"
	"gopkg.in/gomail.v2"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func GenerateOTP() (uint64, string) {
	charset := []byte("123456789")
	code := make([]byte, 7)

	for i := range code {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		code[i] = charset[n.Int64()]
	}

	otpStr := string(code)
	otp, _ := strconv.ParseUint(otpStr, 10, 32)

	return uint64(otp), otpStr
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
	htmlTemplate := `
    <!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Verification</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .container {
            width: 100%;
            max-width: 600px;
            margin: auto;
            background: #ffffff;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #333;
            font-size: 24px;
        }
        p {
            color: #555;
            font-size: 16px;
        }
        .code {
            font-weight: bold;
            font-size: 24px;
            color: #007bff;
            margin: 20px 0;
        }
        .footer {
            text-align: center;
            margin-top: 30px;
            font-size: 12px;
            color: #888;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Please Verify Your Email Address</h1>
        <p>Dear New User,</p>
        <p>Thank you for signing up! To complete your registration, please verify your email address by entering the code below:</p>
        <p class="code">{{OTP}}</p>
        <p>Copy and paste the code into the verification field on our website to proceed.</p>
        <p>If you did not sign up for an account, please disregard this email.</p>
    </div>
</body>
</html>

    `
	htmlBody := strings.Replace(htmlTemplate, "{{OTP}}", otp, 1)
	m.SetBody("text/html", htmlBody)

	// Create dialer
	d := gomail.NewDialer(mailSMTP, mailPort, mailUsername, mailPassword)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
