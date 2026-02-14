package mysqluserrepository

import (
	"database/sql"
	"fmt"
	"q-game-app/entity"
	"q-game-app/pkg/errmsg"
	"q-game-app/pkg/richerror"
	"q-game-app/repository/mysql"
)

func (d *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	var exists bool
	err := d.conn.Conn().QueryRow(`
        SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = ?)
    `, phoneNumber).Scan(&exists)

	if err != nil {
		return false, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return !exists, nil // unique if not exists
}

func (d *DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"

	row := d.conn.Conn().QueryRow(`select id, name, phone_number, password, role from users where phone_number = ?`, phoneNumber)

	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		// TODO - log unexpected error for better observability
		return entity.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return user, nil
}

func (d *DB) RegisterUser(u entity.User) (entity.User, error) {
	res, err := d.conn.Conn().Exec(`
        INSERT INTO users(name, phone_number, password, role)
        VALUES (?, ?, ?, ?)
    `, u.Name, u.PhoneNumber, u.Password, u.Role.String())

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

func (d *DB) GetUserByID(userID uint) (entity.User, error) {

	row := d.conn.Conn().QueryRow(`select id, name, phone_number, password, role from users where id = ?`, userID)
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

func scanUser(scanner mysql.Scanner) (entity.User, error) {
	var user entity.User
	var roleStr string

	err := scanner.Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
		&roleStr,
	)
	if err != nil {
		return entity.User{}, err
	}

	user.Role = entity.MapToRoleEntity(roleStr)

	return user, nil
}
