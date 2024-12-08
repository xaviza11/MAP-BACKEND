package services

import (
	"go-sqlite-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditRailways(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	country := database.Country{
		Name: "TestLand7",
		LAT:  12.34,
		LON:  56.78,
	}

	countryInfo := database.CountryInfo{
		Notes: "Test for editing railways",
	}

	err := CreateCountry(country, countryInfo)
	assert.NoError(t, err)

	var countryID int
	err = database.DB.QueryRow("SELECT id FROM countries WHERE name = ?", country.Name).Scan(&countryID)
	assert.NoError(t, err)

	initialRailways := []database.Railway{
		{Name: "Railway A"},
		{Name: "Railway B"},
	}

	err = EditRailways(countryID, initialRailways)
	assert.NoError(t, err)

	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM railways WHERE country_info_id = ? AND name = ?", countryID, "Railway A").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM railways WHERE country_info_id = ? AND name = ?", countryID, "Railway B").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	updatedRailways := []database.Railway{
		{Name: "Railway C"},
		{Name: "Railway D"},
	}

	err = EditRailways(countryID, updatedRailways)
	assert.NoError(t, err)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM railways WHERE country_info_id = ? AND name = ?", countryID, "Railway C").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM railways WHERE country_info_id = ? AND name = ?", countryID, "Railway D").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM railways WHERE country_info_id = ? AND name = ?", countryID, "Railway A").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM railways WHERE country_info_id = ? AND name = ?", countryID, "Railway B").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}
