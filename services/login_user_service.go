package services

import (
	"errors"
	"go-sqlite-backend/database"
	"go-sqlite-backend/utils"
)

func Login(email, password string) (*database.User, error) {
	var user database.User
	row := database.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
