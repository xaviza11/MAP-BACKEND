package services

import (
	"go-sqlite-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditCountry(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	country := database.Country{
		Name: "TestLand2",
		LAT:  12.34,
		LON:  56.78,
	}

	err := CreateCountry(country, database.CountryInfo{})
	assert.NoError(t, err)

	var countryID int
	err = database.DB.QueryRow("SELECT id FROM countries WHERE name = ?", country.Name).Scan(&countryID)
	assert.NoError(t, err)

	newName := "UpdatedLand"
	newLAT := 98.76
	newLON := 54.32

	err = EditCountry(countryID, newName, newLAT, newLON)
	assert.NoError(t, err)

	var updatedName string
	var updatedLAT, updatedLON float64
	err = database.DB.QueryRow("SELECT name, lat, lon FROM countries WHERE id = ?", countryID).Scan(&updatedName, &updatedLAT, &updatedLON)
	assert.NoError(t, err)

	assert.Equal(t, newName, updatedName)
	assert.Equal(t, newLAT, updatedLAT)
	assert.Equal(t, newLON, updatedLON)
}
