package models

import (
	"backend/db"
	"backend/utils"
	"errors"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLogin struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (username, email, password, created_at)
		VALUES (?, ?, ?, ?)
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Username, u.Email, hashedPassword, u.CreatedAt)
	if err != nil {
		return err
	}
	u.ID, err = result.LastInsertId()
	return err
}

func (u *UserLogin) ValidateCreds() error {
	query := "select id, password from users where email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPwd string
	err := row.Scan(&u.ID, &retrievedPwd)
	if err != nil {
		return err
	}

	pwdIsValid := utils.CheckPasswordHash(u.Password, retrievedPwd)
	if !pwdIsValid {
		return errors.New("Credentials are Invalid")
	}
	return nil
}
