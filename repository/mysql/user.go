package mysql

import (
	"database/sql"
	"fmt"
	"q-game-app/entity"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	var exists bool
	err := d.db.QueryRow(`
        SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = ?)
    `, phoneNumber).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("failed to check phone number uniqueness: %w", err)
	}

	return !exists, nil // unique if not exists
}

func (d *MysqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := d.db.QueryRow(`
        SELECT id, name, phone_number, password
        FROM users
        WHERE phone_number = ?
    `, phoneNumber)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, fmt.Errorf("failed to get user by phone number: %w", err)
	}

	return user, true, nil
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
			return entity.User{}, fmt.Errorf("user not found")
		}
		return entity.User{}, fmt.Errorf("failed to get user by ID: %w", err)
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
