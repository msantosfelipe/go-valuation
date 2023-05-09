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

const (
	stocksBrFilePath  = "resources/input_stocks_br.txt"
	fiisBrFilePath    = "resources/input_fiis_br.txt"
	fiagrosBrFilePath = "resources/input_fiagros_br.txt"
)

func main() {
	var indicatorsProvider = provider.NewProviderInstance()

	processStocks(indicatorsProvider)
	processFiis(indicatorsProvider)
	processFiagros(indicatorsProvider)
}

func processStocks(indicatorsProvider provider.Provider) {
	var brStocks = getFromFile(stocksBrFilePath)
	var wg sync.WaitGroup
	wg.Add(len(brStocks))

	color.HiGreen("Starting to process %v stocks: %v", len(brStocks), brStocks)

	ch := make(chan provider.IndicatorsStock, len(brStocks))
	for _, stock := range brStocks {
		go indicatorsProvider.GetBrStocksIndicators(stock, ch, &wg)
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

func processFiis(indicatorsProvider provider.Provider) {
	var brFiis = getFromFile(fiisBrFilePath)
	var wg sync.WaitGroup
	wg.Add(len(brFiis))

	color.HiGreen("Starting to process %v FIIs: %v", len(brFiis), brFiis)

	ch := make(chan provider.IndicatorsFiis, len(brFiis))
	for _, stock := range brFiis {
		go indicatorsProvider.GetBrFiisIndicators(stock, ch, &wg)
	}

	wg.Wait()
	close(ch)
	color.HiGreen("All FIIs have been processed...")

	for indicators := range ch {
		color.Magenta("\n****************************")
		color.Magenta("\n*** %s ***", indicators.Name)
		indicators.ToString()
	}
}

func processFiagros(indicatorsProvider provider.Provider) {
	var brFiis = getFromFile(fiagrosBrFilePath)
	var wg sync.WaitGroup
	wg.Add(len(brFiis))

	color.HiGreen("Starting to process %v FIIs: %v", len(brFiis), brFiis)

	ch := make(chan provider.IndicatorsFiis, len(brFiis))
	for _, stock := range brFiis {
		go indicatorsProvider.GetBrFiagrosIndicators(stock, ch, &wg)
	}

	wg.Wait()
	close(ch)
	color.HiGreen("All FIIs have been processed...")

	for indicators := range ch {
		color.Magenta("\n****************************")
		color.Magenta("\n*** %s ***", indicators.Name)
		indicators.ToString()
	}
}

func getFromFile(filePath string) []string {
	var brStocks = []string{}

	file, err := os.Open(filePath)
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
