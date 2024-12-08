package services

import (
	"go-sqlite-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditExports(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	country := database.Country{
		Name: "TestLand3",
		LAT:  12.34,
		LON:  56.78,
	}

	countryInfo := database.CountryInfo{
		Notes: "Test for editing exports",
	}

	err := CreateCountry(country, countryInfo)
	assert.NoError(t, err)

	var countryID int
	err = database.DB.QueryRow("SELECT id FROM countries WHERE name = ?", country.Name).Scan(&countryID)
	assert.NoError(t, err)

	initialExports := []database.Export{
		{Name: "Gold", Quantity: 100},
		{Name: "Silver", Quantity: 200},
	}

	err = EditExports(countryID, initialExports)
	assert.NoError(t, err)

	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM exports WHERE country_info_id = ? AND name = ?", countryID, "Gold").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM exports WHERE country_info_id = ? AND name = ?", countryID, "Silver").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	updatedExports := []database.Export{
		{Name: "Gold", Quantity: 500},
		{Name: "Platinum", Quantity: 1000},
	}

	err = EditExports(countryID, updatedExports)
	assert.NoError(t, err)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM exports WHERE country_info_id = ? AND name = ?", countryID, "Gold").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM exports WHERE country_info_id = ? AND name = ?", countryID, "Platinum").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM exports WHERE country_info_id = ? AND name = ?", countryID, "Silver").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}
