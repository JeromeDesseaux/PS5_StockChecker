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

	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0")
	request.Header.Set("Accept", "image/webp,*/*")
	//request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept-Language", "fr,fr-FR;q=0.8,en-US;q=0.5,en;q=0.3")

	response, err := client.Do(request)
	if response.StatusCode == 403 {
		return errors.New(fmt.Sprintf("Impossible d'accéder à la page %s", config.URL))
	}

	for _, assertion := range config.Assertions {
		if assertion.Selector != "" {
			doc, err := goquery.NewDocumentFromReader(response.Body)
			defer response.Body.Close()
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
