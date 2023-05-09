package provider

import (
	"fmt"
	"sync"
)

type ProviderResponse struct {
	Name        string
	ActualPrice float64
	VPA         float64
	LPA         float64
	ActualDY    string
	PastYearDY  float64
}

func (p *ProviderResponse) ToString() string {
	return fmt.Sprintf("Actual price: R$%v\n\nVPA: %v\nLPA: %v\nActual DY: %s\nPast Year DY: %v\n\n",
		p.ActualPrice,
		p.VPA,
		p.LPA,
		p.ActualDY,
		p.PastYearDY)
}

type Provider interface {
	GetBrStockIndicators(stock string, ch chan ProviderResponse, wg *sync.WaitGroup)
}
