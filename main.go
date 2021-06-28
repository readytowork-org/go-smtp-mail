package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

func main() {

	// Sender data.
	from := "apikey"
	password := "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"

	// Receiver email address.
	to := []string{
		"prajwol.kc@readytowork.jp",
	}

	// smtp server configuration.
	smtpHost := "smtp.sendgrid.net"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("From: prajwol.kc@pcmgmt.edu.np\nTo: prajwol.kc@readytowork.jp\nSubject: This is a test subject\n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "23423 Puneet Singh 23",
		Message: "Thi 23 4324 s is a test message in a HTML template",
	})
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
