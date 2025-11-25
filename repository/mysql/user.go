package mysql

import (
	"fmt"
	"q-game-app/entity"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	var exists bool
	err := d.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = ?)`, phoneNumber).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check phone number uniqueness: %w", err)
	}
	return !exists, nil
}

func (d *MysqlDB) RegisterUser(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number, password) values(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't insert user: %w", err)
	}
	id, err := res.LastInsertId()
	u.ID = uint(id)
	return u, nil

}
