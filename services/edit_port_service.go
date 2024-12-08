package services

import (
	"go-sqlite-backend/database"
)

func EditPorts(countryID int, ports []database.Port) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM ports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, port := range ports {
		_, err = tx.Exec("INSERT INTO ports(country_info_id, name) VALUES(?, ?)", countryID, port.Name)
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
