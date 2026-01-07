package mysql

import (
	"database/sql"
	"fmt"
	"q-game-app/entity"
	"q-game-app/pkg/errmsg"
	"q-game-app/pkg/richerror"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	var exists bool
	err := d.db.QueryRow(`
        SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = ?)
    `, phoneNumber).Scan(&exists)

	if err != nil {
		return false, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return !exists, nil // unique if not exists
}

func (d *MysqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`
        SELECT id, name, phone_number, password
        FROM users
        WHERE phone_number = ?
    `, phoneNumber)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}
		return entity.User{}, richerror.New(op).WithMessage(errmsg.ErrorMsgCantScanQueryResult)

	}

	return user, nil
}

func (d *MysqlDB) RegisterUser(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`
        INSERT INTO users(name, phone_number, password)
        VALUES (?, ?, ?)
    `, u.Name, u.PhoneNumber, u.Password)

	if err != nil {
		return entity.User{}, fmt.Errorf("failed to insert user: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to fetch inserted user ID: %w", err)
	}

	u.ID = uint(id)
	return u, nil
}

func (d *MysqlDB) GetUserByID(userID uint) (entity.User, error) {
	row := d.db.QueryRow(`
        SELECT id, name, phone_number, password
        FROM users
        WHERE id = ?
    `, userID)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New("mysql.GetUserByID")
		}
		return entity.User{}, richerror.New("mysql.GetUserByID")
	}

	return user, nil
}

// ----------------------------------------------------
// HELPER
// ----------------------------------------------------

func scanUser(row *sql.Row) (entity.User, error) {
	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
	)
	return user, err
}
