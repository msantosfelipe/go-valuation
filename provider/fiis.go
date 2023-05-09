package provider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fiiFromStatusInvest(fii string, doc *goquery.Document) IndicatorsFiis {
	vpc, pvp := getFiiIndicators(doc)
	return IndicatorsFiis{
		Name:        fii,
		ActualPrice: getFiiActualPrice(doc),
		VPC:         vpc,
		PVP:         pvp,
		ActualDY:    getFiiActualDY(doc),
		PastYearDY:  getFiiPastYearDY(doc),
	}
}

func getFiiActualPrice(doc *goquery.Document) float64 {
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

func getFiiIndicators(doc *goquery.Document) (float64, float64) {
	var vpc, pvp float64
	var err error

	doc.Find(`.top-info`).Each(func(i int, s *goquery.Selection) {
		s.Find("div").Each(func(i int, s1 *goquery.Selection) {
			if s1.Find("h3").Text() == "Val. patrimonial p/cota" {
				v := strings.Replace(strings.TrimSpace(s1.Find("strong").Text()), ",", ".", 1)
				if vpc, err = strconv.ParseFloat(v, 64); err != nil {
					fmt.Println(err)
				}
				return
			}

			if s1.Find("h3").Text() == "P/VP" {
				v := strings.Replace(strings.TrimSpace(s1.Find("strong").Text()), ",", ".", 1)
				if pvp, err = strconv.ParseFloat(v, 64); err != nil {
					fmt.Println(err)
				}
				return
			}
		})
	})

	return vpc, pvp
}

func getFiiActualDY(doc *goquery.Document) string {
	var actualDY string

	selection := doc.Find(`[title="Dividend Yield com base nos últimos 12 meses"]`).First()
	stringValue := selection.Find("strong").Text()
	if stringValue == "" {
		return "0%"
	}

	actualDY = stringValue + "%"

	return actualDY
}

func getFiiPastYearDY(doc *goquery.Document) float64 {
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
