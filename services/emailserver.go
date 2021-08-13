package services

import "fmt"

type Email struct {
	To      string
	Subject string
	Body    string
}

//NewEmail is a initializer or constructor for type Email
func NewEmail(to, subject, body string) Email {
	return Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func SendEmail(email Email) bool {
	fmt.Println("Sending email to ", email.To)
	fmt.Println("Email Subject: ", email.Subject)
	fmt.Println("Email Body: \n", email.Body)

	return true
}
