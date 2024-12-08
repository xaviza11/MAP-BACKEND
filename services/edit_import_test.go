package services

import (
	"go-sqlite-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditImports(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	country := database.Country{
		Name: "TestLand5",
		LAT:  12.34,
		LON:  56.78,
	}

	countryInfo := database.CountryInfo{
		Notes: "Test for editing imports",
	}

	err := CreateCountry(country, countryInfo)
	assert.NoError(t, err)

	var countryID int
	err = database.DB.QueryRow("SELECT id FROM countries WHERE name = ?", country.Name).Scan(&countryID)
	assert.NoError(t, err)

	initialImports := []database.Import{
		{Name: "Oil", Quantity: 300},
		{Name: "Steel", Quantity: 150},
	}

	err = EditImports(countryID, initialImports)
	assert.NoError(t, err)

	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM imports WHERE country_info_id = ? AND name = ?", countryID, "Oil").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM imports WHERE country_info_id = ? AND name = ?", countryID, "Steel").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	updatedImports := []database.Import{
		{Name: "Natural Gas", Quantity: 500},
		{Name: "Copper", Quantity: 200},
	}

	err = EditImports(countryID, updatedImports)
	assert.NoError(t, err)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM imports WHERE country_info_id = ? AND name = ?", countryID, "Natural Gas").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM imports WHERE country_info_id = ? AND name = ?", countryID, "Copper").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM imports WHERE country_info_id = ? AND name = ?", countryID, "Oil").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM imports WHERE country_info_id = ? AND name = ?", countryID, "Steel").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}
