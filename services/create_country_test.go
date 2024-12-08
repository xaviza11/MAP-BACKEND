package services

import (
	"testing"

	"go-sqlite-backend/database"

	"github.com/stretchr/testify/assert"
)

func TestCreateCountry(t *testing.T) {
	database.InitializeTestDB()
	defer database.CloseTestDB()

	country := database.Country{
		Name: "TestLand",
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

	var count int

	db := database.DB

	err = db.QueryRow("SELECT COUNT(*) FROM countries WHERE name = ?", country.Name).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.QueryRow("SELECT COUNT(*) FROM country_info WHERE notes = ?", countryInfo.Notes).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.QueryRow("SELECT COUNT(*) FROM factions WHERE name = ?", "Faction A").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.QueryRow("SELECT COUNT(*) FROM exports WHERE name = ?", "Gold").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.QueryRow("SELECT COUNT(*) FROM imports WHERE name = ?", "Oil").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.QueryRow("SELECT COUNT(*) FROM ports WHERE name = ?", "Port A").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	err = db.QueryRow("SELECT COUNT(*) FROM railways WHERE name = ?", "Railway A").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}
