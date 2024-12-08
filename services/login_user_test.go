package services

import (
	"go-sqlite-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	user := database.User{
		Name:     "TestUser8",
		Email:    "testuser2@example2.com",
		Password: "ABCdsdfsdfe1234123asd",
	}

	err := CreateUser(user)
	assert.NoError(t, err)

	loggedInUser, err := Login(user.Email, user.Password)
	assert.NoError(t, err)
	assert.NotNil(t, loggedInUser)
	assert.Equal(t, user.Name, loggedInUser.Name)
	assert.Equal(t, user.Email, loggedInUser.Email)

	_, err = Login(user.Email, "wrongPassword123")
	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())

	_, err = Login("nonexistent@example.com", "somepassword")
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}
