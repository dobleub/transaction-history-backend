package email

import (
	"crypto/tls"
	"log"

	gomail "gopkg.in/mail.v2"

	"github.com/dobleub/transaction-history-backend/internal/config"
)

func SendEmail(env *config.EmailConfig, emailFrom string, emailTo string, subject string, body string, bodyText string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if bodyText != "" {
		m.AddAlternative("text/plain", bodyText)
	}

	d := gomail.NewDialer(env.SMTPHost, env.SMTPPort, env.SMTPUsername, env.SMTPPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: env.STMPSecure}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		// panic(err)
		return err
	}

	return nil
}
