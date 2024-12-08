package services

import (
	"go-sqlite-backend/database"
	"go-sqlite-backend/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	user := database.User{
		Name:     "TestUser",
		Email:    "testuser@example.com",
		Password: "TestPassword123",
	}

	err := CreateUser(user)
	assert.NoError(t, err)

	var count int
	db := database.DB

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ?", user.Name).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", user.Email).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	var storedHashedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE email = ?", user.Email).Scan(&storedHashedPassword)
	assert.NoError(t, err)

	assert.NotEqual(t, user.Password, storedHashedPassword)

	err = utils.ComparePassword(storedHashedPassword, user.Password)
	assert.NoError(t, err)
}
