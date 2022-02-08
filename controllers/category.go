package controllers

import (
	"errors"
	"strings"

	"github.com/idobalul/dark-web-scraping/models"
)

type Results struct {
	General             float64
	Hacking             float64
	Crypto              float64
	Market              float64
	MaybeScams          float64
	IllegalAdultContent float64
}

var (
	cryptoKeyWords  = []string{"bitcion", "generator"}
	hackingKeyWords = []string{"hacker", "hacking"}
	marketKeyWords  = []string{"market", "shop", "buy", "payment"}
	scamKeyWords    = []string{"scam", "get $"}
	adultKeyWords   = []string{"c.p", "loli", "sex", "s3x", "porn", "child", "teen"}
)

// Categorize checks the category of the paste
// and extracts the presentage analysis.
func Categorize(pastes []models.Paste) (Results, error) {
	// Initiating the vars for the analysis.
	var (
		total           = len(pastes)
		precentPerPaste = 100 / float64(total)
		general         = 0
		hacking         = 0
		crypto          = 0
		market          = 0
		maybeScams      = 0
		adultContent    = 0
	)

	// If there are no pastes return error.
	if len(pastes) == 0 {
		return Results{}, errors.New("No pastes to categorize")
	}

	// For every paste in the pastes slice check the category.
	for _, paste := range pastes {
		// Check the category of the paste.
		category := checkCategory(paste.Title, paste.Content)

		switch category {
		case "adultContent":
			adultContent += 1
		case "crypto":
			crypto += 1
		case "hacking":
			hacking += 1
		case "market":
			market += 1
		case "scam":
			maybeScams += 1
		default:
			general += 1
		}
	}

	results := Results{
		General:             float64(general) * precentPerPaste,
		Hacking:             float64(hacking) * precentPerPaste,
		Crypto:              float64(crypto) * precentPerPaste,
		Market:              float64(market) * precentPerPaste,
		MaybeScams:          float64(maybeScams) * precentPerPaste,
		IllegalAdultContent: float64(adultContent) * precentPerPaste,
	}

	return results, nil
}

// checkCategory checks the category of the paste
// with the keywords arrays.
func checkCategory(title string, content []string) string {
	titleForCheck := strings.ToLower(title)
	contentForCheck := strings.ToLower(strings.Join(content, "\n"))

	// Checks for adult content keywords in the title and content.
	for _, word := range adultKeyWords {
		if strings.Contains(titleForCheck, word) {
			return "adultContent"
		}
		if strings.Contains(contentForCheck, word) {
			return "adultContent"
		}
	}

	// Checks for hacking content keywords in the title and content.
	for _, word := range hackingKeyWords {
		if strings.Contains(titleForCheck, word) {
			return "hacking"
		}
		if strings.Contains(contentForCheck, word) {
			return "hacking"
		}
	}

	// Checks for crypto content keywords in the title and content.
	for _, word := range cryptoKeyWords {
		if strings.Contains(titleForCheck, word) {
			return "crypto"
		}
		if strings.Contains(contentForCheck, word) {
			return "crypto"
		}
	}

	// Checks for marketing content keywords in the title and content.
	for _, word := range marketKeyWords {
		if strings.Contains(titleForCheck, word) {
			return "market"
		}
		if strings.Contains(contentForCheck, word) {
			return "market"
		}
	}

	// Checks for scam content keywords in the title and content.
	for _, word := range scamKeyWords {
		if strings.Contains(titleForCheck, word) {
			return "scam"
		}
		if strings.Contains(contentForCheck, word) {
			return "scam"
		}
	}

	// If there is no category return general.
	return "general"
}
