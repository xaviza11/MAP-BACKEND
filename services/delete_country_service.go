package services

import (
	"go-sqlite-backend/database"
)

func DeleteCountry(countryID int) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM factions WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM exports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM imports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM ports WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM railways WHERE country_info_id IN (SELECT id FROM country_info WHERE country_id = ?)", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM country_info WHERE country_id = ?", countryID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM countries WHERE id = ?", countryID)
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
