package db

import (
	"log"
	"strings"

	"github.com/idobalul/dark-web-scraping/models"
)

// GetPastes returns all the pastes from the database.
func GetPastes() ([]models.Paste, error) {
	var pastes []models.Paste

	rows, err := db.Query("SELECT title, content, author, date FROM pastes")
	if err != nil {
		log.Println("Error getting pastes from database:", err)
	}

	for rows.Next() {
		var paste models.Paste

		// content holds the string from the content column.
		var content string

		err := rows.Scan(&paste.Title, &content, &paste.Author, &paste.Date)
		if err != nil {
			log.Println("Error scanning paste from database:", err)
			continue
		}

		// Split the content string into an array of strings.
		paste.Content = strings.Split(content, "\n")

		pastes = append(pastes, paste)
	}

	// If any uncaught error can happen, it catches it.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pastes, nil
}
