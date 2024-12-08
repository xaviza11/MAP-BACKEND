package services

import (
	"go-sqlite-backend/database"
)

func EditExports(countryID int, exports []database.Export) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM exports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, export := range exports {
		_, err = tx.Exec("INSERT INTO exports(country_info_id, name, quantity) VALUES(?, ?, ?)", countryID, export.Name, export.Quantity)
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
