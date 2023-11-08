package utils

import (
	"fmt"
	"net/smtp"
)

type Mail struct {
	Host     string
	Port     string
	From     string
	Password string
	To       []string
	Subject  string
	Message  string
}

func (m *Mail) SendMail() error {
	auth := smtp.PlainAuth("", m.From, m.Password, m.Host)
	msg := []byte(fmt.Sprintf("To: %s\r\n"+"Subject: %s\r\n"+"\r\n"+"%s\r\n", m.To[0], m.Subject, m.Message))

	if err := smtp.SendMail(m.Host+":"+m.Port, auth, m.From, m.To, msg); err != nil {
		return err
	}

	return nil
}

func MailMessage(uuid string) string {
	message := fmt.Sprintf("以下のリンクをクリックしてメールアドレスを認証してください。\nhttp://localhost:8080/auth/mail?uuid=%s\nひとりでウミガメのスープへのアカウント登録に身に覚えがない場合は無視してください。", uuid)

	return message
}
