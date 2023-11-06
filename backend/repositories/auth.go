package repositories

import (
	"database/sql"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/myerrors"
)

func RegisterUser(db *sql.DB, auth *models.Auth) error {
	const sqlStr = `
	insert into users (email, password_hash, activate_flag, created_at) values 
	(?, ?, false, now());
	`

	if _, err := db.Exec(sqlStr, auth.Email, auth.Hash); err != nil {
		err = myerrors.InsertDataFailed.Wrap(err, "internal server error")
		return err
	}

	return nil
}
