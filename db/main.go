package db

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

// ConnectToDB opens a connection to the database
// and returns a pointer to the database.
func ConnectToDB() {
	// Opens a connection with the database.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}

	// Checking if the db connection works.
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// Create the table if it doesn't exist.
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pastes (id INT NOT NULL AUTO_INCREMENT, title VARCHAR(255) NOT NULL, author VARCHAR(255) NOT NULL, content LONGTEXT NOT NULL, date VARCHAR(128) NOT NULL, PRIMARY KEY (id), UNIQUE INDEX date_UNIQUE (date ASC) VISIBLE);`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database.")
}
