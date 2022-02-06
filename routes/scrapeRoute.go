package routes

import (
	"net/http"

	"github.com/idobalul/dark-web-scraping/controllers"
	"github.com/idobalul/dark-web-scraping/db"

	"github.com/gin-gonic/gin"
)

// scrapeRouter initialize the scrape router.
func scrapeRouter(rg *gin.RouterGroup) {
	rg.GET("/scrape", func(c *gin.Context) {
		// Scraping the dark web.
		controllers.Scrape()

		// Getting the pastes from the database.
		pastes, err := db.GetPastes()
		// If there is an error, send to the client status 500.
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// If there is no error, send the pastes to the client.
		c.JSON(http.StatusOK, gin.H{"pastes": pastes})
	})
}
