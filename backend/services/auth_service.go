package services

import (
	"database/sql"
	"net/mail"
	"os"

	"github.com/google/uuid"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/myerrors"
	"github.com/yoshi-zen/sea-turtle/backend/repositories"
	"github.com/yoshi-zen/sea-turtle/backend/utils"
)

func RegisterUserService(db *sql.DB, auth *models.Auth) error {
	if _, err := mail.ParseAddress(auth.Email); err != nil {
		err = myerrors.EmailInvalid.Wrap(err, "email invalid")
		return err
	}

	hash, err := utils.PasswordEncrypt(auth.Password)
	if err != nil {
		err = myerrors.EncryptFailed.Wrap(err, "internal server error")
		return err
	}
	auth.Hash = hash

	uuidObj, err := uuid.NewUUID()
	if err != nil {
		err = myerrors.GenUUIDFailed.Wrap(err, "internal server error")
		return err
	}
	mailAuthUuid := uuidObj.String()
	auth.Uuid = mailAuthUuid
	mailMessage := utils.MailMessage(mailAuthUuid)
	myMail := utils.Mail{
		Host:     "smtp.gmail.com",
		Port:     "587",
		From:     os.Getenv("EMAIL_ADDRESS"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		To:       []string{auth.Email},
		Subject:  "メールアドレスの認証",
		Message:  mailMessage,
	}
	if err := myMail.SendMail(); err != nil {
		err = myerrors.SendMailFailed.Wrap(err, "internal server error")
		return err
	}

	if err = repositories.RegisterUser(db, auth); err != nil {
		return err
	}

	return nil
}

func MailCheckService(db *sql.DB, uuid string) error {
	err := repositories.UpdateActivate(db, uuid)
	if err != nil {
		return err
	}

	return nil
}
