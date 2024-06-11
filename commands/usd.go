package commands

import (
	"fmt"
	"github.com/EugenSleptsov/triban/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type BankRate struct {
	BankName string
	Buying   float64
	Selling  float64
}

type UsdCommand struct {
	lastFetchTime time.Time
	previousRates map[string]BankRate
	currentRates  map[string]BankRate
	mutex         sync.Mutex
}

func (d *UsdCommand) Execute(args []string) string {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the rates were fetched within the last 10 minutes
	if time.Since(d.lastFetchTime) < 10*time.Minute {
		return d.formatRates(d.currentRates, d.previousRates, d.lastFetchTime)
	}

	url := "https://kur.doviz.com/kapalicarsi/amerikan-dolari"

	// Getting content
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to get the URL: %v", err)
		return "Ошибка получения курсов валюты USD"
	}
	defer resp.Body.Close()

	// Checking status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to get the URL: %s", resp.Status)
		return "Ошибка получения курсов валюты USD"
	}

	// Parsing the HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Failed to parse the HTML: %v", err)
		return "Ошибка получения курсов валюты USD"
	}

	// Extracting the rates
	newRates := make(map[string]BankRate)
	doc.Find(".value-table .table.table-narrow.sortable table tbody tr").Each(func(i int, s *goquery.Selection) {
		bankName := strings.TrimSpace(s.Find("td:nth-child(1)").Text())
		buyingStr := strings.TrimSpace(s.Find("td:nth-child(2)").Text())
		sellingStr := strings.TrimSpace(s.Find("td:nth-child(3)").Text())

		// Parse the buying and selling rates from string to float64
		buying, err := strconv.ParseFloat(strings.Replace(buyingStr, ",", ".", 1), 64)
		if err != nil {
			log.Printf("Failed to parse buying rate for %s: %v", bankName, err)
			return
		}
		selling, err := strconv.ParseFloat(strings.Replace(sellingStr, ",", ".", 1), 64)
		if err != nil {
			log.Printf("Failed to parse selling rate for %s: %v", bankName, err)
			return
		}

		newRates[bankName] = BankRate{
			BankName: bankName,
			Buying:   buying,
			Selling:  selling,
		}
	})

	if len(newRates) == 0 {
		log.Printf("Empty bank rates")
		return "Ошибка получения курсов валюты USD"
	}

	// Update the cache
	d.previousRates = d.currentRates
	d.currentRates = newRates
	d.lastFetchTime = time.Now()

	return d.formatRates(d.currentRates, d.previousRates, d.lastFetchTime)
}
func (d *UsdCommand) formatRates(rates map[string]BankRate, lastRates map[string]BankRate, lastFetchTime time.Time) string {
	var result string
	result += "Курсы валюты USD на момент " + lastFetchTime.Format("15:04:05") + ":\n\n"

	sortedKeys := sortedKeys(rates)

	bestBuy, bestSell := bestBuySellRate(rates)

	for _, bankName := range sortedKeys {
		rate := rates[bankName]
		buyingRate := utils.FloatToString(rate.Buying)
		sellingRate := utils.FloatToString(rate.Selling)

		// Bold the best buy rate
		if rate.Buying == bestBuy.Buying {
			buyingRate = "❗" + buyingRate + "❗"
		}

		// Bold the best sell rate
		if rate.Selling == bestSell.Selling {
			sellingRate = "❗" + sellingRate + "❗"
		}

		if lastRate, ok := lastRates[bankName]; ok {
			// Calculate changes
			buyingChange := rate.Buying - lastRate.Buying
			sellingChange := rate.Selling - lastRate.Selling

			result += fmt.Sprintf(
				"%s\nПокупка: %s (%+.4f) | Продажа: %s (%+.4f)\n\n",
				bankName,
				buyingRate, buyingChange,
				sellingRate, sellingChange,
			)
		} else {
			// No previous info, show only current rates
			result += fmt.Sprintf(
				"%s\nПокупка: %s | Продажа: %s\n\n",
				bankName,
				buyingRate,
				sellingRate,
			)
		}
	}
	return result
}

func (d *UsdCommand) Description() string {
	return "Информация о курсах валюты USD в разных банках"
}

func sortedKeys(unsortedMap map[string]BankRate) []string {
	keys := make([]string, 0, len(unsortedMap))
	for key := range unsortedMap {
		keys = append(keys, key)
	}
	utils.SortStrings(keys)
	return keys
}

func bestBuySellRate(rates map[string]BankRate) (bestBuy BankRate, bestSell BankRate) {
	// Initialize bestBuy and bestSell with extreme values to ensure proper comparison
	bestBuy = BankRate{Buying: -1.0, Selling: 0.0}
	bestSell = BankRate{Buying: 0.0, Selling: 1e9}
	for _, rate := range rates {
		// this banks should not be checked
		if map[string]bool{"Harem": true, "Odacı": true, "Kapalıçarşı": true, "Merkez Bankası": true, "Venüs": true, "Altınkaynak": true}[rate.BankName] {
			continue
		}

		if rate.Buying > bestBuy.Buying {
			bestBuy = rate
		}
		if rate.Selling < bestSell.Selling {
			bestSell = rate
		}
	}
	return
}
