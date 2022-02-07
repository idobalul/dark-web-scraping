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
	// Initiate the collector
	collector := colly.NewCollector()

	// Use proxy to connect to TOR
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:9050")
	if err != nil {
		log.Fatal(err)
	}
	collector.SetProxyFunc(rp)

	// for every paste block on the page collect the title, content, author and date
	collector.OnHTML(".col-sm-12", func(container *colly.HTMLElement) {
		var content []string

		// get the title
		title := strings.Trim(container.ChildText("h4"), " \n")

		// get the content
		container.ForEach("li", func(i int, element *colly.HTMLElement) {
			content = append(content, strings.Trim(element.Text, " \n"))
		})

		// Only if the paste has a tittle continue normally
		if title != "" {
			rawInfo := strings.Trim(container.ChildText(".pre-footer"), " \n")
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

	collector.Visit("http://strongerw2ise74v3duebgsvug4mehyhlpa7f6kfwnas7zofs3kov7yd.onion/all")
	collector.Wait()
	db.AddPastes(pastes)
}
