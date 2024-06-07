package commands

import (
	"fmt"
	"github.com/EugenSleptsov/triban/utils"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type BankRate struct {
	BankName string
	Buying   float64
	Selling  float64
}

type UsdCommand struct{}

func (d UsdCommand) Execute(args []string) string {
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

	var BankRates []BankRate

	// Extracting the <tr> elements
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

		BankRates = append(BankRates, BankRate{
			BankName: bankName,
			Buying:   buying,
			Selling:  selling,
		})
	})

	if len(BankRates) == 0 {
		log.Printf("Empty bank rates")
		return "Ошибка получения курсов валюты USD"
	}

	var result string
	for _, rate := range BankRates {
		result += fmt.Sprintf("%s\nПокупка: %s\nПродажа: %s\n\n", rate.BankName, utils.FloatToString(rate.Buying), utils.FloatToString(rate.Selling))
	}

	return result
}

func (d UsdCommand) Description() string {
	return "Информация о курсах валюты USD в разных банках"
}
