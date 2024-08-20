package function

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendMail(toMail string, token int) {
	smtpHost := os.Getenv("SMTP_HOST")
	emailSender := os.Getenv("EMAIL_SENDER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	tmpl := `Hello, your reset token is: <b>{{.Token}}</b>`

	data := struct {
		Token int
	}{
		Token: token,
	}

	t, err := template.New("email").Parse(tmpl)
	if err != nil {
		fmt.Printf("error parsing template: %v", err)
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		fmt.Printf("error executing template: %v", err)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", emailSender)
	m.SetHeader("To", toMail)
	m.SetHeader("Subject", "Reset Password")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(smtpHost, 587, emailSender, emailPassword)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
