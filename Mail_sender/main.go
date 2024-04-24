package main

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
)

func generateOTP(length int) string {
	const charset = "0123456789"
	seededRand := rand.New(rand.NewSource(rand.Int63()))
	otp := make([]byte, length)
	for i := range otp {
		idx := seededRand.Intn(len(charset))
		otp[i] = charset[idx]
	}
	return string(otp)
}

func sendOTP(email, otp string) error {
	// Sender's email configuration
	sender := "mahindra7396@gmail.com"
	password := "Harshi@14"
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// Recipient's email address
	to := []string{email}

	// Message setup
	subject := "OTP for Registration"
	body := "Your OTP is: " + otp

	// Connect to the SMTP server
	smtpAddress := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	auth := smtp.PlainAuth("", sender, password, smtpHost)
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", sender, strings.Join(to, ","), subject, body)

	// Send email
	err := smtp.SendMail(smtpAddress, auth, sender, to, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	email := "mahindragajjala2@gmail.com" // Replace with the recipient's email address

	// Generate OTP
	otp := generateOTP(6) // Generate a 6-digit OTP

	// Send OTP to email
	err := sendOTP(email, otp)
	if err != nil {
		fmt.Println("Error sending OTP:", err)
		return
	}

	fmt.Println("OTP sent successfully to", email)
}
