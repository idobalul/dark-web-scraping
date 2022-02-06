package db

import (
	"log"
	"strings"

	"github.com/idobalul/dark-web-scraping/models"
)

// AddPastes receives a slice of pastes and adds them to the database.
func AddPastes(pastes []models.Paste) {
	for _, paste := range pastes {
		// formating the content to be inserted.
		content := strings.Join(paste.Content, "\n")

		// Inserting the paste into the database.
		_, err := db.Exec("INSERT INTO pastes (title, content, author, date) VALUES (?, ?, ?, ?)", paste.Title, content, paste.Author, paste.Date)
		if err != nil {
			log.Println("Error inserting paste into database:", err)
		}
	}
}
