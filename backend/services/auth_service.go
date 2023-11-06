package services

import (
	"database/sql"
	"net/mail"
	"os"

	"github.com/google/uuid"

	"github.com/yoshi-zen/sea-turtle/backend/controllers"
	"github.com/yoshi-zen/sea-turtle/backend/myerrors"
	"github.com/yoshi-zen/sea-turtle/backend/repositories"
	"github.com/yoshi-zen/sea-turtle/backend/utils"
)

func RegisterUserService(db *sql.DB, auth *controllers.Auth) error {
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

	repositories.RegisterUser(db, auth)

	uuidObj, err := uuid.NewUUID()
	if err != nil {
		err = myerrors.GenUUIDFailed.Wrap(err, "internal server error")
		return err
	}

	uuid := uuidObj.String()

	/*
		ToDo: メール本文考える、認証するAPIリクエストを送るリンクの送信
	*/
	myMail := utils.Mail{
		from:     os.Getenv("EMAIL_ADDRESS"),
		password: os.Getenv("EMAIL_PASSWORD"),
		to:       auth.email,
		subject:  mailSubject,
		message:  mailMessage,
	}
	if err := myMail.SendMail(); err != nil {
		err = myerrors.SendMailFailed.Wrap(err, "internal server error")
		return err
	}

	return nil
}
