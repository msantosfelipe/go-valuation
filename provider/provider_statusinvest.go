package provider

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/msantosfelipe/go-valuation/config"
)

type providerStatusInvest struct{}

func NewProviderInstance() Provider {
	return &providerStatusInvest{}
}

func (p *providerStatusInvest) GetBrStockIndicators(stock string) ProviderResponse {
	baseUrl := fmt.Sprintf("%s/acoes/%s", config.ENV.UrlStockStatusInvest, stock)
	doc := getProviderPage(baseUrl)
	return findFromFundamentus(doc)
}

func getProviderPage(baseUrl string) *goquery.Document {
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7,es-MX;q=0.6,es;q=0.5")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`)
	req.Header.Set("sec-ch-ua-mobile", "70")
	req.Header.Set("sec-ch-ua-platform", "macOS")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func findFromFundamentus(doc *goquery.Document) ProviderResponse {
	return ProviderResponse{
		ActualPrice: getActualPrice(doc),
		VPA:         getVPA(doc),
		LPA:         getLPA(doc),
		ActualDY:    getActualDY(doc),
		PastYearDY:  getPastYearDY(doc),
	}
}

func getActualPrice(doc *goquery.Document) float64 {
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

func getVPA(doc *goquery.Document) float64 {
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

func getLPA(doc *goquery.Document) float64 {
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

func getActualDY(doc *goquery.Document) string {
	var actualDY string

	selection := doc.Find(`[title="Dividend Yield com base nos últimos 12 meses"]`).First()
	stringValue := selection.Find("strong").Text()
	if stringValue == "" {
		return "0%"
	}

	actualDY = stringValue + "%"

	return actualDY
}

func getPastYearDY(doc *goquery.Document) float64 {
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
