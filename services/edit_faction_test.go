package services

import (
	"go-sqlite-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditFactions(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	country := database.Country{
		Name: "TestLand4",
		LAT:  12.34,
		LON:  56.78,
	}

	countryInfo := database.CountryInfo{
		Notes: "Test for editing factions",
	}

	err := CreateCountry(country, countryInfo)
	assert.NoError(t, err)

	var countryID int
	err = database.DB.QueryRow("SELECT id FROM countries WHERE name = ?", country.Name).Scan(&countryID)
	assert.NoError(t, err)

	initialFactions := []database.Faction{
		{Name: "Faction1", Support: "70"},
		{Name: "Faction2", Support: "30"},
	}

	err = EditFactions(countryID, initialFactions)
	assert.NoError(t, err)

	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM factions WHERE country_info_id = ? AND name = ?", countryID, "Faction1").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM factions WHERE country_info_id = ? AND name = ?", countryID, "Faction2").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	updatedFactions := []database.Faction{
		{Name: "Faction3", Support: "50"},
		{Name: "Faction4", Support: "50"},
	}

	err = EditFactions(countryID, updatedFactions)
	assert.NoError(t, err)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM factions WHERE country_info_id = ? AND name = ?", countryID, "Faction3").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM factions WHERE country_info_id = ? AND name = ?", countryID, "Faction4").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM factions WHERE country_info_id = ? AND name = ?", countryID, "Faction1").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM factions WHERE country_info_id = ? AND name = ?", countryID, "Faction2").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}
