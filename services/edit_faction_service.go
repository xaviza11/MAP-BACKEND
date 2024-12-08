package services

import (
	"go-sqlite-backend/database"
)

func EditFactions(countryID int, factions []database.Faction) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM factions WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, faction := range factions {
		_, err = tx.Exec("INSERT INTO factions(country_info_id, name, support) VALUES(?, ?, ?)", countryID, faction.Name, faction.Support)
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
