package models

import (
	"errors"
	"go-event-api/db"
	"go-event-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?,?)`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = id

	return nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string

	err := row.Scan(&u.ID, &hashedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, hashedPassword)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}
	return nil

}
