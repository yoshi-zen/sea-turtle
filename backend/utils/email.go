package utils

import (
	"fmt"
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

func MailMessage(uuid string) string {
	message := fmt.Sprintf(`以下のリンクをクリックしてメールアドレスを認証してください。
	http://localhost:8080/auth/mail?uuid=%s
	ひとりでウミガメのスープへのアカウント登録に身に覚えがない場合は無視してください。`, uuid)

	return message
}
