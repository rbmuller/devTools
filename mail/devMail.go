package devMail

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail() {

	// Set up the email message
	m := gomail.NewMessage()
	m.SetHeader("From", "robson.muller@valor-software.com")
	m.SetHeader("To", "robson.muller@valor-software.com")
	m.SetHeader("Subject", "Test Email")
	m.SetBody("text/html", "Hello <b>there</b>!")

	// Set up the email server
	d := gomail.NewDialer("smtp.gmail.com", 587, "your-email@gmail.com", "your-password")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}
