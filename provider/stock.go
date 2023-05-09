package provider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func stockFromStatusInvest(stock string, doc *goquery.Document) IndicatorsStock {
	return IndicatorsStock{
		Name:        stock,
		ActualPrice: getStockActualPrice(doc),
		VPA:         getStockVPA(doc),
		LPA:         getStockLPA(doc),
		ActualDY:    getStockActualDY(doc),
		PastYearDY:  getStockPastYearDY(doc),
	}
}

func getStockActualPrice(doc *goquery.Document) float64 {
	selection := doc.Find(`[title="Valor atual do ativo"]`).First()
	stringValue := selection.Find("strong").Text()
	if stringValue == "" {
		return 0
	}

	v := strings.Replace(strings.TrimSpace(stringValue), ",", ".", 1)
	vpa, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Println(err)
	}

	return vpa
}

func getStockVPA(doc *goquery.Document) float64 {
	selection := doc.Find(`[title="Indica qual o valor patrimonial de uma ação."]`).First()
	stringValue := selection.Find("strong").Text()
	if stringValue == "" {
		return 0
	}

	v := strings.Replace(strings.TrimSpace(stringValue), ",", ".", 1)
	vpa, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Println(err)
	}

	return vpa
}

func getStockLPA(doc *goquery.Document) float64 {
	selection := doc.Find(`[title="Indicar se a empresa é ou não lucrativa. Se este número estiver negativo, a empresa está com margens baixas, acumulando prejuízos."]`).First()

	stringValue := selection.Find("strong").Text()
	if stringValue == "" {
		return 0
	}

	v := strings.Replace(strings.TrimSpace(stringValue), ",", ".", 1)
	lpa, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Println(err)
	}

	return lpa
}

func getStockActualDY(doc *goquery.Document) string {
	var actualDY string

	selection := doc.Find(`[title="Dividend Yield com base nos últimos 12 meses"]`).First()
	stringValue := selection.Find("strong").Text()
	if stringValue == "" {
		return "0%"
	}

	actualDY = stringValue + "%"

	return actualDY
}

func getStockPastYearDY(doc *goquery.Document) float64 {
	selection := doc.Find(`[title="Soma total de proventos distribuídos nos últimos 12 meses"]`).First()
	stringValue := selection.Find(".sub-value").Text()
	if stringValue == "" {
		return 0
	}

	valueSplitted := strings.Split(stringValue, "R$")
	if len(valueSplitted) != 2 {
		return 0
	}

	v := strings.Replace(strings.TrimSpace(valueSplitted[1]), ",", ".", 1)
	pastYearDY, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Println(err)
	}

	return pastYearDY
}
