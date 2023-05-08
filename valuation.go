package main

import (
	"math"

	"github.com/msantosfelipe/go-valuation/config"
)

// Benjamin Grahan formula - 'value = √ (22,5 x LPA x VPA)'
func CalculateGrahamStockValue(lpa, vpa float64) float64 {
	value := math.Sqrt((22.5 * lpa * vpa))
	return toFixed(value, 2)
}

// Decio Bazin formula - 'value = (PastYearDY x 100)/6'
func CalculateBazinStockValue(pastYearDY float64) float64 {
	value := (pastYearDY * 100) / config.ENV.BazinDividendPercentage
	return toFixed(value, 2)
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
