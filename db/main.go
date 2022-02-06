package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// ConnectToDB opens a connection to the database
// and returns a pointer to the database.
func ConnectToDB() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "a6c(tM$=E6!vU5W1ID;.",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "scraper",
	}

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
