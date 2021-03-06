package controllers

import (
	"log"
	"strings"

	"github.com/idobalul/dark-web-scraping/db"
	"github.com/idobalul/dark-web-scraping/models"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

var pastes []models.Paste

// Scrape is the main function for the scraper.
// It scrapes the pastebin and stores the pastes in the database.
func Scrape() {
	log.Println("Scraping the dark web...")

	// Clearing the pastes array before scraping.
	pastes = []models.Paste{}

	// Initiate the collector
	c := colly.NewCollector()

	// Use proxy to connect to TOR
	rp, err := proxy.RoundRobinProxySwitcher("socks5://torproxy:9050")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// for every paste block on the page collect the title, content, author and date
	c.OnHTML(".col-sm-12", func(h *colly.HTMLElement) {
		var content []string

		// get the title
		title := strings.Trim(h.ChildText("h4"), " \n")

		// get the content
		h.ForEach("li", func(i int, e *colly.HTMLElement) {
			content = append(content, strings.Trim(e.Text, " \n"))
		})

		// Only if the paste has a tittle continue normally
		if title != "" {
			rawInfo := strings.Trim(h.ChildText(".pre-footer"), " \n")
			info := strings.Split(rawInfo, "by")
			authorAndDate := strings.Split(info[1], "at")

			// extract author and date from the raw info
			author := authorAndDate[0]
			date := strings.SplitAfter(authorAndDate[1], "UTC")[0]

			// add the paste to the pastes slice
			pastes = append(pastes, models.Paste{
				Title:   title,
				Content: content,
				Author:  author,
				Date:    date,
			})
		}
	})

	err = c.Visit("http://strongerw2ise74v3duebgsvug4mehyhlpa7f6kfwnas7zofs3kov7yd.onion/all")
	if err != nil {
		log.Println(err)
	}
	c.Wait()
	db.AddPastes(pastes)
}
