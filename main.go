package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
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

// writeToCSV writes the stock data to a CSV file
func writeToCSV(stocks []Stock, filename string) error {
	//Create the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Symbol", "Company Name", "Stock Price", "Change"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header to CSV: %w", err)
	}

	// Write data
	for _, stock := range stocks {
		row := []string{stock.Symbol, stock.CompanyName, stock.StockPrice, stock.Change}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing row to CSV: %w", err)
		}
	}

	return nil
}

func main() {
	// URL of the page to scrape
	scrapeURL := "https://stockanalysis.com/list/faang-stocks/"
	
	// Scrape the stock data
	stocks, err := scrapeStocks(scrapeURL)
	if err != nil {
		fmt.Printf("Error scraping data: %s\n", err)
		return
	}

	// Write the scraped data to a CSV file
	err = writeToCSV(stocks, "faang_stocks.csv")
	if err != nil {
		fmt.Printf("Error writing to CSV: %s\n", err)
		return
	}

	fmt.Println("Data written to faang_stocks.csv successfully!")
}

