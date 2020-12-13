package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JeromeDesseaux/scraper/notifications"
	"github.com/JeromeDesseaux/scraper/parsers"
	"github.com/JeromeDesseaux/scraper/scraper"
)

func main() {
	config, err := parsers.ReadConfig("config/ps5.json")
	if err != nil {
		fmt.Println("Impossible de lire le fichier donné. Vérifiez s'il existe et s'il est correctement formé.\n", err)
		os.Exit(1)
	}
	for _, c := range config {
		err = scraper.CheckURL(&c)
		if err != nil {
			if strings.Contains(err.Error(), "Impossible d'accéder") {
				title := fmt.Sprintf("%s : Changement de statut", c.Website)
				notifications.ShowNotification(title, "Une ou plusieurs règles ne fonctionne plus sur ce site. Allez voir, votre article est peut-être en stock !")
			}
		}
	}
}
