package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PublicGetJson(url string, v interface{}) error {
	resp, err := http.Get(url)
	log.Println("GET", url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()



	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Print("Response body: ", string(body))

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		return errors.New(fmt.Sprintf("Got status code %d", resp.StatusCode))
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

type ServerTimeResponse struct {
	Time int64 `json:"time"`
}

func GetServerTime() (int64, error) {
	var serverTimeResponse ServerTimeResponse
	url := BuildURL("/time")
	err := PublicGetJson(url, &serverTimeResponse)
	if err != nil {
		return serverTimeResponse.Time, err
	}
	if serverTimeResponse.Time == 0 {
		return serverTimeResponse.Time, errors.New("server time was 0")
	}
	return serverTimeResponse.Time, nil
}

type TickerResponse struct {
	Market Market `json:"market"`
	Bid    FiatAmount `json:"bid"`
	Ask    FiatAmount `json:"ask"`
	Spread FiatAmount `json:"spread"`
}

// GetTicket gets a market ticker for a specific market.
func GetTicker(market Market) (TickerResponse, error) {
	var tickerResponse TickerResponse
	url := BuildURL("/v2/markets/%s/ticker", market)
	err := PublicGetJson(url, &tickerResponse)
	if tickerResponse.Market == "" {
		tickerResponse.Market = market
	}
	return tickerResponse, err
}
