package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Stock represents a single stock entry
type Stock struct {
	Symbol      string
	CompanyName string
	StockPrice  string
	Change      string
}

// scrapeStocks scrapes FAANG stock data from the given URL
func scrapeStocks(url string) ([]Stock, error) {
	var stocks []Stock
	c := colly.NewCollector()

	// Define the scraping logic for the main table
	c.OnHTML("#main-table", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			// Extract text from each column
			stock := Stock{
				Symbol:      row.ChildText(".sym"),
				CompanyName: row.ChildText(".slw"),
				StockPrice:  row.ChildText("td:nth-of-type(4)"),
				Change:      row.ChildText("td:nth-of-type(5)"),
			}
			stocks = append(stocks, stock)
		})
	})

	// Log when a request is made
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %s\n", r.URL)
	})

	// Log any errors that occur during scraping
	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	// Start the scraping process
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}
