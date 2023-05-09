package provider

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
)

type IndicatorsStock struct {
	Name        string
	ActualPrice float64
	VPA         float64 // PREÇO/VALOR PATRIMONIAL P/AÇÃO
	LPA         float64 // LUCRO/VALOR POR AÇÃO
	ActualDY    string
	PastYearDY  float64
}

func (p *IndicatorsStock) ToString() string {
	return fmt.Sprintf("Actual price: R$%v\n\nVPA: %v\nLPA: %v\nActual DY: %s\nPast Year DY: %v\n\n",
		p.ActualPrice,
		p.VPA,
		p.LPA,
		p.ActualDY,
		p.PastYearDY,
	)
}

type IndicatorsFiis struct {
	Name        string
	ActualPrice float64
	VPC         float64 // VALOR PATRIMONIAL P/COTA
	PVP         float64 // PREÇO/VALOR PATRIMONIAL
	ActualDY    string
	PastYearDY  float64
}

func (p *IndicatorsFiis) ToString() {
	color.Cyan("Actual price: R$%v", p.ActualPrice)
	color.Cyan("VPC: %v", p.VPC)
	if p.PVP > 1 {
		color.Red("PVP: %v", p.PVP)
	} else {
		color.Green("PVP: %v", p.PVP)
	}
	color.Cyan("nActual DY: %v", p.ActualDY)
	color.Cyan("Past Year DY: %v", p.PastYearDY)
}

type Provider interface {
	GetBrStocksIndicators(stock string, ch chan IndicatorsStock, wg *sync.WaitGroup)
	GetBrFiisIndicators(stock string, ch chan IndicatorsFiis, wg *sync.WaitGroup)
	GetBrFiagrosIndicators(stock string, ch chan IndicatorsFiis, wg *sync.WaitGroup)
}
