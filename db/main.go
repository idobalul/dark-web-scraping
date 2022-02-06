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
		log.Fatal(err)
	}

	// Checking if the db connection works.
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")
}
