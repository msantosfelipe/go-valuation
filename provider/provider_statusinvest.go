package provider

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/msantosfelipe/go-valuation/config"
)

type providerStatusInvest struct{}

func NewProviderInstance() Provider {
	return &providerStatusInvest{}
}

func (p *providerStatusInvest) GetBrStocksIndicators(stock string, ch chan IndicatorsStock, wg *sync.WaitGroup) {
	defer wg.Done()

	baseUrl := fmt.Sprintf("%s/acoes/%s", config.ENV.UrlStockInvest, stock)
	doc := getProviderPage(baseUrl)
	ch <- stockFromStatusInvest(stock, doc)
	color.HiGreen("Stock %s processed...", stock)
}

func (p *providerStatusInvest) GetBrFiisIndicators(stock string, ch chan IndicatorsFiis, wg *sync.WaitGroup) {
	defer wg.Done()

	baseUrl := fmt.Sprintf("%s/fundos-imobiliarios/%s", config.ENV.UrlStockInvest, stock)
	doc := getProviderPage(baseUrl)
	_ = doc
	ch <- fiiFromStatusInvest(stock, doc)
	color.HiGreen("FII %s processed...", stock)
}

func (p *providerStatusInvest) GetBrFiagrosIndicators(stock string, ch chan IndicatorsFiis, wg *sync.WaitGroup) {
	defer wg.Done()

	baseUrl := fmt.Sprintf("%s/fiagros/%s", config.ENV.UrlStockInvest, stock)
	doc := getProviderPage(baseUrl)
	_ = doc
	ch <- fiiFromStatusInvest(stock, doc)
	color.HiGreen("FII %s processed...", stock)
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
