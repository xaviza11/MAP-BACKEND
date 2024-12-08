package services

import (
	"go-sqlite-backend/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteCountry(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	country := database.Country{
		Name: "TestLand1",
		LAT:  12.34,
		LON:  56.78,
	}

	countryInfo := database.CountryInfo{
		Notes: "A fictional country for testing",
		Factions: []database.Faction{
			{Name: "Faction A", Support: "70"},
			{Name: "Faction B", Support: "30"},
		},
		Exports: []database.Export{
			{Name: "Gold", Quantity: 100},
			{Name: "Silver", Quantity: 200},
		},
		Imports: []database.Import{
			{Name: "Oil", Quantity: 300},
			{Name: "Steel", Quantity: 150},
		},
		Ports: []database.Port{
			{Name: "Port A"},
			{Name: "Port B"},
		},
		Railway: []database.Railway{
			{Name: "Railway A"},
			{Name: "Railway B"},
		},
	}

	err := CreateCountry(country, countryInfo)
	assert.NoError(t, err)

	var countryID int
	err = database.DB.QueryRow("SELECT id FROM countries WHERE name = ?", country.Name).Scan(&countryID)
	assert.NoError(t, err)

	err = DeleteCountry(countryID)
	assert.NoError(t, err)

	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM countries WHERE id = ?", countryID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM country_info WHERE country_id = ?", countryID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM factions WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM exports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM imports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM ports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)

	err = database.DB.QueryRow("SELECT COUNT(*) FROM railways WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}
