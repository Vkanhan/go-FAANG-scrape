package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

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
