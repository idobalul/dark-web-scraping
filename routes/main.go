package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/idobalul/dark-web-scraping/controllers"
	"github.com/idobalul/dark-web-scraping/db"
)

var server = gin.Default()

// getRoutes sets up the routes for the server.
func getRoutes() {
	main := server.Group("/")
	scrapeRouter(main)
}

// Run starts the server.
func Run() {
	db.ConnectToDB()

	// Scrapint for the first time when the server starts.
	controllers.Scrape()
	// Set the scraping interval to run concurrently.
	ticker := time.NewTicker(time.Minute * 2)
	go func() {
		for {
			<-ticker.C
			// Scraping the dark web.
			controllers.Scrape()

		}
	}()

	//Setting up the routes and starting to listen on port 8080.
	getRoutes()
	server.Run(":8080")
}
