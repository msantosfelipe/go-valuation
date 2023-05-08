package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/msantosfelipe/go-valuation/provider"
)

func main() {
	var provider = provider.NewProviderInstance()
	var brStocks = getStocksFromFile()

	for _, stock := range brStocks {
		color.Magenta("\n*** %s ***", stock)
		indicators := provider.GetBrStockIndicators(stock)
		color.Cyan(indicators.ToString())

		grahan := CalculateGrahamStockValue(indicators.LPA, indicators.VPA)
		bazin := CalculateBazinStockValue(indicators.PastYearDY)
		color.White("Graham's intrinsic value: %v", grahan)
		color.White("Bazin's fair price: %v", bazin)
	}

	// save all data into output file
}

func getStocksFromFile() []string {
	var brStocks = []string{}

	file, err := os.Open("resources/input_stocks_br.txt")
	if err != nil {
		log.Fatal("error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		brStocks = append(brStocks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
	}

	return brStocks
}
