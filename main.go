package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/msantosfelipe/go-valuation/provider"
)

func main() {
	var indicatorsProvider = provider.NewProviderInstance()
	var brStocks = getStocksFromFile()

	var wg sync.WaitGroup
	wg.Add(len(brStocks))

	color.HiGreen("Starting to process %v stocks: %v", len(brStocks), brStocks)

	ch := make(chan provider.ProviderResponse, len(brStocks))
	for _, stock := range brStocks {
		go indicatorsProvider.GetBrStockIndicators(stock, ch, &wg)
	}

	wg.Wait()
	close(ch)
	color.HiGreen("All stocks have been processed...")

	for indicators := range ch {
		color.Magenta("\n****************************")
		color.Magenta("\n*** %s ***", indicators.Name)
		color.Cyan(indicators.ToString())

		grahan := CalculateGrahamStockValue(indicators.LPA, indicators.VPA)
		bazin := CalculateBazinStockValue(indicators.PastYearDY)
		color.White("Graham's intrinsic value: %v", grahan)
		color.White("Bazin's fair price: %v", bazin)
	}
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
