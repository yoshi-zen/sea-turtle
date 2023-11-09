package repositories

import (
	"database/sql"

	"github.com/yoshi-zen/sea-turtle/backend/models"
	"github.com/yoshi-zen/sea-turtle/backend/myerrors"
)

func RegisterUser(db *sql.DB, auth *models.Auth) error {
	const sqlStr = `
	insert into users (email, password_hash, uuid, activate_flag, created_at) values 
	(?, ?, ?, false, now());
	`

	if _, err := db.Exec(sqlStr, auth.Email, auth.Hash, auth.Uuid); err != nil {
		err = myerrors.InsertDataFailed.Wrap(err, "internal server error")
		return err
	}

	return nil
}

func UpdateActivate(db *sql.DB, uuid string) error {
	const sqlStr = `
	update users
	set activate_flag = true
	where uuid = ?;
	`

	result, err := db.Exec(sqlStr, uuid)
	if err != nil {
		err = myerrors.UpdateDataFailed.Wrap(err, "internal server error")
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		err = myerrors.Unknown.Wrap(err, "internal server error")
		return err
	}
	if num == 0 {
		err = myerrors.NoData.Wrap(NotFound, "not found")
		return err
	}

	return nil
}
