package services

import (
	"go-sqlite-backend/database"
)

func CreateCountry(country database.Country, countryInfo database.CountryInfo) error {

	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO countries(name, lat, lon) VALUES(?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(country.Name, country.LAT, country.LON)
	if err != nil {
		tx.Rollback()
		return err
	}

	countryID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	stmtInfo, err := tx.Prepare("INSERT INTO country_info(country_id, notes) VALUES(?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmtInfo.Close()

	_, err = stmtInfo.Exec(countryID, countryInfo.Notes)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, faction := range countryInfo.Factions {
		stmtFaction, err := tx.Prepare("INSERT INTO factions(country_info_id, name, support) VALUES(?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmtFaction.Exec(countryID, faction.Name, faction.Support)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, export := range countryInfo.Exports {
		stmtExport, err := tx.Prepare("INSERT INTO exports(country_info_id, name, quantity) VALUES(?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmtExport.Exec(countryID, export.Name, export.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, importItem := range countryInfo.Imports {
		stmtImport, err := tx.Prepare("INSERT INTO imports(country_info_id, name, quantity) VALUES(?, ?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmtImport.Exec(countryID, importItem.Name, importItem.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, port := range countryInfo.Ports {
		stmtPort, err := tx.Prepare("INSERT INTO ports(country_info_id, name) VALUES(?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmtPort.Exec(countryID, port.Name)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, railway := range countryInfo.Railway {
		stmtRailway, err := tx.Prepare("INSERT INTO railways(country_info_id, name) VALUES(?, ?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmtRailway.Exec(countryID, railway.Name)
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
