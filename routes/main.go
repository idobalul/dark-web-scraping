package routes

import (
	"github.com/gin-gonic/gin"
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
	getRoutes()
	server.Run(":8080")
}
