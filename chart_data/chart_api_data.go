package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type APIResp struct {
	BPI        map[string]float64 `json:"bpi"`
	Disclaimer string             `json:"disclaimer"`
	Time       struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
	} `json:"time"`
}

// getRESTCall is a more generic implementation of http request for reusability.
// Gets a http response from RESTful api (or any json formatted response). Body
// of the response is put into a []byte and then unmarshaled into a generic
// interface which is returned to the calling function. Referenced:
// https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9
func getRESTCall(url string) APIResp {
	var respData APIResp

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return respData
	}

	// For control over HTTP client headers, redirect policy, and other
	// settings, create a Client A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client Do sends an HTTP request and returns an
	// HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("http client do error: ", err)
		return respData
	}

	// Callers should close resp.Body when done reading from it Defer the
	// closing of the body
	defer resp.Body.Close()

	// Put body of the response into a []byte
	body, err := ioutil.ReadAll(resp.Body)

	// Unmarshal to APIResp struct and return if ok
	err = json.Unmarshal(body, &respData)
	if err != nil {
		log.Fatal("Json unmarshal error: ", err)
		return respData
	}

	return respData
}

// buildPriceMap parses relevent pricing data and creates a map from it.
func buildPriceMap(resp APIResp, priceMap map[int]float64) {
	priceDataRaw := resp.BPI
	for k, v := range priceDataRaw {
		i := parseDateToInt(k)
		priceMap[i] = v
	}
}

func parseDateToInt(sDate string) int {
	trimmedString := strings.Replace(sDate, "-", "", -1)
	i, err := strconv.Atoi(trimmedString)
	if err != nil {
		log.Fatal("Error converting date to int: ", err)
		return 0
	}
	return i
}

// xyBuilder returns some random x, y points.
func xyBuilder(n int, dates []int, prices []float64) plotter.XYs {
	pts := make(plotter.XYs, n)
	index := 0
	for i, _ := range dates {
		// TODO make X axis the actual dates, right now its just a sequence from
		// 0 to n
		pts[index].X = float64(i)
		pts[index].Y = prices[i]
		index++
	}
	return pts
}

func toOrderedArray(priceMap map[int]float64) ([]int, []float64) {
	keys := make([]int, 0)
	for k := range priceMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	values := make([]float64, 0)
	for _, i := range keys {
		values = append(values, priceMap[i])
	}

	return keys, values
}

func main() {
	// Custom start to current date
	t := time.Now().Local()
	date := t.Format("2006-01-02")
	start := "2016-01-01"
	input := fmt.Sprintf("https://api.coindesk.com/v1/bpi/historical/close.json?start=%s&end=%s", start, date)
	dataTimePeriod := fmt.Sprintf(" (%s to %s)", start, date)
	// Past 30 days
	// input := "https://api.coindesk.com/v1/bpi/historical/close.json?"
	// dataTimePeriod := " (31 Days)"

	apiJson := getRESTCall(input)
	var priceMap map[int]float64
	priceMap = make(map[int]float64)
	buildPriceMap(apiJson, priceMap)
	dates, prices := toOrderedArray(priceMap)
	sizeOfDataSet := len(dates)

	// Create chart
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Bitcoin Price Chart" + dataTimePeriod
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Price (USD)"
	// Draw a grid behind the data
	p.Add(plotter.NewGrid())
	p.HideX()

	err = plotutil.AddLinePoints(p,
		"Bitcoin", xyBuilder(sizeOfDataSet, dates, prices))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(15*vg.Inch, 15*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
