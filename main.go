package main

import (
	"encoding/csv"
	"fmt"
	"os"
	

	"github.com/gocolly/colly/v2"
)

func main() {

	scrapeURL := "https://stockanalysis.com/list/faang-stocks/"

	c := colly.NewCollector()

	// Slice to hold the scraped data
	var data [][]string

	c.OnHTML("#main-table", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			symbol := row.ChildText(".sym")
			companyName := row.ChildText(".slw")
			stockPrice := row.ChildText("td:nth-of-type(4)")
			change := row.ChildText("td:nth-of-type(5)")

			// Append the data to the slice
			data = append(data, []string{symbol, companyName, stockPrice, change})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	c.OnScraped(func(r *colly.Response) {
		// Create a CSV file
		file, err := os.Create("faang_stocks.csv")
		if err != nil {
			fmt.Println("Error creating CSV file:", err)
			return
		}
		defer file.Close()

		// Create a CSV writer
		writer := csv.NewWriter(file)
		defer writer.Flush()

		// Write header
		header := []string{"Symbol", "Company Name", "Stock Price", "Change"}
		if err := writer.Write(header); err != nil {
			fmt.Println("Error writing header to CSV:", err)
			return
		}

		// Write data to CSV
		for _, row := range data {
			if err := writer.Write(row); err != nil {
				fmt.Println("Error writing row to CSV:", err)
				return
			}
		}

		fmt.Println("Data written to faang_stocks.csv successfully!")
	})

	c.Visit(scrapeURL)
}




