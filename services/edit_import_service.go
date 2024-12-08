package services

import (
	"go-sqlite-backend/database"
)

func EditImports(countryID int, imports []database.Import) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM imports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, importItem := range imports {
		_, err = tx.Exec("INSERT INTO imports(country_info_id, name, quantity) VALUES(?, ?, ?)", countryID, importItem.Name, importItem.Quantity)
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
