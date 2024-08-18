package main

import (
	"fmt"
)

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
