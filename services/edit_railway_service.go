package services

import (
	"go-sqlite-backend/database"
)

func EditRailways(countryID int, railways []database.Railway) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM railways WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, railway := range railways {
		_, err = tx.Exec("INSERT INTO railways(country_info_id, name) VALUES(?, ?)", countryID, railway.Name)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
