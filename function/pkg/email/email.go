package email

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"github.com/elulcao/stori-function/pkg/processor"

	gomail "gopkg.in/gomail.v2"
)

//go:embed assets/*
var assets embed.FS

type EmailData struct {
	Transaction *processor.Transaction
	Sender      string
}

// sendEmail is a helper function that sends the email.
func sendEmail(data *processor.Transaction) (err error) {
	sender := "email@example.com"
	ed := &EmailData{
		Transaction: data,
		Sender:      sender,
	}
	data.Owner = "John Doe User"

	tmpl, err := template.ParseFS(assets, "assets/email.html")
	if err != nil {
		return fmt.Errorf("unable to parse template: %w", err)
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, ed)
	if err != nil {
		return fmt.Errorf("unable to execute template: %w", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Monthly Transaction Summary")
	m.SetBody("text/html", tpl.String())

	d := gomail.Dialer{Host: "mailhog", Port: 1025} // Mock SMTP server
	if err = d.DialAndSend(m); err != nil {
		return fmt.Errorf("email message not created: %w", err)
	}

	return nil
}

// SendEmail is the exported entry point for the Cloud Function.
func SendEmail(data *processor.Transaction) (err error) {
	err = sendEmail(data)
	if err != nil {
		return fmt.Errorf("unable to create email: %w", err)
	}

	return nil
}
