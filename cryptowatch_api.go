package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Total amounts held
const (
	BtcAmount = 0
	LtcAmount = 0
	ZecAmount = 0
)

// CoinPrice - type for returning REST call data colleted from cryptowatch.ch.
// Use for converting josn to types: https://mholt.github.io/json-to-go/
type CoinPrice struct {
	Result struct {
		Price float64 `json:"price"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

// MarketStatus - type that get all market status
// API uses: https://api.cryptowat.ch/exchanges
type MarketStatus struct {
	Result []struct {
		ID     int    `json:"id"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Route  string `json:"route"`
		Active bool   `json:"active"`
	} `json:"result"`
	Allowance struct {
		Cost      int   `json:"cost"`
		Remaining int64 `json:"remaining"`
	} `json:"allowance"`
}

// getCoinPrice Creates a http api url to call the cryptowatch api.
func getCoinPrice(coin string) float64 {
	// Builds a url to send via http
	url := fmt.Sprintf("https://api.cryptowat.ch/markets/bitfinex/%s/price", coin)

	// Call api
	resp := getRESTCall(url)
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var priceData CoinPrice
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&priceData); err != nil {
		log.Println(err)
	}

	return priceData.Result.Price
}

// getRESTCall A more generic implementation of http request for reusability
// Referenced: https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9
func getRESTCall(url string) *http.Response {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	return resp
}

func main() {
	btcPrice := getCoinPrice("btcusd")
	fmt.Println("BTC Price: ", btcPrice)

	ltcPrice := getCoinPrice("ltcusd")
	fmt.Println("LTC Price: ", ltcPrice)

	zecPrice := getCoinPrice("zecusd")
	fmt.Println("ZEC Price: ", zecPrice)

	totalV := (btcPrice * BtcAmount) + (ltcPrice * LtcAmount) + (ltcPrice * LtcAmount)
	fmt.Println("Total Value:", totalV)
}
