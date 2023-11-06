package utils

import (
	"net/smtp"
)

type Mail struct {
	from     string
	password string
	to       []string
	subject  string
	message  string
}

func (m *Mail) SendMail() error {
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", m.from, m.password, smtpHost)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, m.from, m.to, []byte(m.subject+m.message)); err != nil {
		return err
	}

	return nil
}
