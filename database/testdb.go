package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var TestDB *sql.DB

func InitializeTestDB() {
	var err error
	TestDB, err = sql.Open("sqlite3", "../test.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
	DB = TestDB
}

func CloseTestDB() {
	if TestDB != nil {
		TestDB.Close()
	}
}

func createTables() {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT UNIQUE,
        email TEXT UNIQUE,
		password TEXT
    );
    `
	_, err := TestDB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createCountryTableSQL := `
    CREATE TABLE IF NOT EXISTS countries (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT UNIQUE,
        lat REAL,
        lon REAL
    );
    `
	_, err = TestDB.Exec(createCountryTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createCountryInfoTableSQL := `
    CREATE TABLE IF NOT EXISTS country_info (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        country_id INTEGER,
        notes TEXT,
        FOREIGN KEY(country_id) REFERENCES countries(id)
    );
    `
	_, err = TestDB.Exec(createCountryInfoTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createFactionTableSQL := `
    CREATE TABLE IF NOT EXISTS factions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        country_info_id INTEGER,
        name TEXT,
        support TEXT,
        FOREIGN KEY(country_info_id) REFERENCES country_info(id)
    );
    `
	_, err = TestDB.Exec(createFactionTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createExportTableSQL := `
    CREATE TABLE IF NOT EXISTS exports (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        country_info_id INTEGER,
        name TEXT,
        quantity INTEGER,
        FOREIGN KEY(country_info_id) REFERENCES country_info(id)
    );
    `
	_, err = TestDB.Exec(createExportTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createImportTableSQL := `
    CREATE TABLE IF NOT EXISTS imports (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        country_info_id INTEGER,
        name TEXT,
        quantity INTEGER,
        FOREIGN KEY(country_info_id) REFERENCES country_info(id)
    );
    `
	_, err = TestDB.Exec(createImportTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createPortTableSQL := `
    CREATE TABLE IF NOT EXISTS ports (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        country_info_id INTEGER,
        name TEXT,
        FOREIGN KEY(country_info_id) REFERENCES country_info(id)
    );
    `
	_, err = TestDB.Exec(createPortTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createRailwayTableSQL := `
    CREATE TABLE IF NOT EXISTS railways (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        country_info_id INTEGER,
        name TEXT,
        FOREIGN KEY(country_info_id) REFERENCES country_info(id)
    );
    `
	_, err = TestDB.Exec(createRailwayTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
