package services

import (
	"go-sqlite-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditPorts(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	country := database.Country{
		Name: "TestLand6",
		LAT:  12.34,
		LON:  56.78,
	}

	countryInfo := database.CountryInfo{
		Notes: "Test for editing ports",
	}

	err := CreateCountry(country, countryInfo)
	assert.NoError(t, err)

	var countryID int
	err = database.DB.QueryRow("SELECT id FROM countries WHERE name = ?", country.Name).Scan(&countryID)
	assert.NoError(t, err)

	initialPorts := []database.Port{
		{Name: "Port A"},
		{Name: "Port B"},
	}

	err = EditPorts(countryID, initialPorts)
	assert.NoError(t, err)

	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM ports WHERE country_info_id = ? AND name = ?", countryID, "Port A").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM ports WHERE country_info_id = ? AND name = ?", countryID, "Port B").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	updatedPorts := []database.Port{
		{Name: "Port C"},
		{Name: "Port D"},
	}

	err = EditPorts(countryID, updatedPorts)
	assert.NoError(t, err)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM ports WHERE country_info_id = ? AND name = ?", countryID, "Port C").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM ports WHERE country_info_id = ? AND name = ?", countryID, "Port D").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM ports WHERE country_info_id = ? AND name = ?", countryID, "Port A").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM ports WHERE country_info_id = ? AND name = ?", countryID, "Port B").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}
