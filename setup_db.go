package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDB() {
	// Open the SQLite database file
	db, err := sql.Open("sqlite3", "./wine.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SQL statements to create the tables
	tableCreationQueries := map[string]string{
		"Wine": `CREATE TABLE IF NOT EXISTS Wine(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name VARCHAR(255) NOT NULL,
            year INTEGER NOT NULL,
            country VARCHAR(255) NOT NULL,
            type VARCHAR(255) NOT NULL,
            description TEXT NOT NULL,
            imageURL VARCHAR(255) NOT NULL,
            volume DECIMAL(10, 2) NOT NULL,
            volAlc DECIMAL(4, 2) NOT NULL,
			UNIQUE (name, year, volume, country)
        );`,
		"Flavour": `CREATE TABLE IF NOT EXISTS Flavour(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name VARCHAR(255) NOT NULL
        );`,
		"FitsTo": `CREATE TABLE IF NOT EXISTS FitsTo(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            description VARCHAR(255) NOT NULL
        );`,
		"Wine_Flavour": `CREATE TABLE IF NOT EXISTS Wine_Flavour(
            wine_id INTEGER,
            flavour_id INTEGER,
            PRIMARY KEY (wine_id, flavour_id),
            FOREIGN KEY (wine_id) REFERENCES Wine(id) ON DELETE CASCADE,
            FOREIGN KEY (flavour_id) REFERENCES Flavour(id) ON DELETE CASCADE
        );`,
		"Wine_FitsTo": `CREATE TABLE IF NOT EXISTS Wine_FitsTo(
            wine_id INTEGER,
            fitsTo_id INTEGER,
            PRIMARY KEY (wine_id, fitsTo_id),
            FOREIGN KEY (wine_id) REFERENCES Wine(id) ON DELETE CASCADE,
            FOREIGN KEY (fitsTo_id) REFERENCES FitsTo(id) ON DELETE CASCADE
        );`,
		"FavoriteWines": `CREATE TABLE IF NOT EXISTS FavoriteWines(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            wine_id INTEGER NOT NULL,
            FOREIGN KEY (wine_id) REFERENCES Wine(id),
            UNIQUE (user_id, wine_id)
        );`,
		"Supermarkets": `CREATE TABLE IF NOT EXISTS Supermarkets(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name VARCHAR(255) NOT NULL,
            street VARCHAR(255) NOT NULL,
            postal_code VARCHAR(20) NOT NULL,
            city VARCHAR(255) NOT NULL,
            house_number VARCHAR(10) NOT NULL
        );`,
		"WineSupermarkets": `CREATE TABLE IF NOT EXISTS WineSupermarkets(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            wine_id INTEGER NOT NULL,
            supermarket_id INTEGER NOT NULL,
            FOREIGN KEY (wine_id) REFERENCES Wine(id),
            FOREIGN KEY (supermarket_id) REFERENCES Supermarkets(id),
            UNIQUE (wine_id, supermarket_id)
        );`,
		"WineFactOTD": `CREATE TABLE IF NOT EXISTS WineFactOTD(
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            fact TEXT NOT NULL
        );`,
	}

	// Execute each table creation query
	for tableName, query := range tableCreationQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error creating table %s: %v", tableName, err)
		} else {
			log.Printf("Table %s created successfully!", tableName)
		}
	}
}
