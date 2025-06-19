package utils

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SendMail(email, token string) {
	message := gomail.NewMessage()
	message.SetHeader("From", "aadesh.kumar@synlabs.io")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Verify your account")

	verifyURL := fmt.Sprintf("http://localhost:8000/verify?token=%s", token)
	body := fmt.Sprintf("Click to verify your email: <a href='%s'>Verify</a>", verifyURL)
	message.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "aadesh.kumar@synlabs.io", "alyrccclqmemptja")
	if err := d.DialAndSend(message); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email sent successfully")
	}
}
