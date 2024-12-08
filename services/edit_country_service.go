package services

import (
	"go-sqlite-backend/database"
)

func EditCountry(countryID int, name string, lat float64, lon float64) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE countries SET name = ?, lat = ?, lon = ? WHERE id = ?", name, lat, lon, countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
