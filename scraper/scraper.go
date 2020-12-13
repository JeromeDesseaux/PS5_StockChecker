package scraper

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/JeromeDesseaux/scraper/parsers"
	"github.com/PuerkitoBio/goquery"
)

// CheckURL verifies the given URL according to the user tests provided in the config file
func CheckURL(config *parsers.WebsiteConfig) error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
		Jar:     jar,
	}

	request, err := http.NewRequest("GET", config.URL, nil)
	if err != nil {
		return err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	request.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)

	response, err := client.Do(request)
	if response.StatusCode == 403 {
		return errors.New(fmt.Sprintf("Impossible d'accéder à la page %s", config.URL))
	}

	for _, assertion := range config.Assertions {
		if assertion.Selector != "" {
			doc, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				return err
			}
			selectorText := doc.Find(assertion.Selector).First().Text()
			for _, text := range assertion.Contains {
				if !strings.Contains(selectorText, text) {
					return errors.New("Le selecteur ne contient pas le texte recherché")
				}
			}
		}
		if assertion.Status != 0 {
			if assertion.Status != response.StatusCode {
				return errors.New(fmt.Sprintf("Mauvais code HTTP %d (attendu : %d)", response.StatusCode, assertion.Status))
			}
		}
	}

	return nil
}
