package services

import (
	"go-sqlite-backend/database"
	"go-sqlite-backend/utils"
)

func CreateUser(user database.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	stmt, err := database.DB.Prepare("INSERT INTO users(name, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, hashedPassword)
	return err
}
