package main

// BASED ON https://www.scrapingbee.com/blog/web-scraping-go/

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("si3.bcentral.cl"),
	)

	yen_found := false

	c.OnHTML("li.list-group-item", func(e *colly.HTMLElement) {
		i := 0
		temp_string := ""
		e.ForEach("label", func(_ int, el *colly.HTMLElement) {
			if i%2 == 0 {
				temp_string = el.Text
			} else {
				if !yen_found {
					writer.Write([]string{
						temp_string,
						el.Text,
					})
					writer.Flush()
					if strings.Contains(temp_string, "Yen") {
						yen_found = true
					}
				}
			}
			i++
		})

	})

	c.Visit("https://si3.bcentral.cl/Indicadoressiete/secure/Indicadoresdiarios.aspx")
}
