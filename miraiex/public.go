package miraiex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (c *Client) PublicGetJson(url string, v interface{}) error {
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
	code := strconv.Itoa(resp.StatusCode)
	log.Print(resp.Header)
	log.Print("Response code", code, "with body:", string(body))

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

func (c *Client) GetServerTime() (int64, error) {
	var serverTimeResponse ServerTimeResponse
	url := BuildURL("/time")
	err := c.PublicGetJson(url, &serverTimeResponse)
	if err != nil {
		return serverTimeResponse.Time, err
	}
	if serverTimeResponse.Time == 0 {
		return serverTimeResponse.Time, errors.New("server time was 0")
	}
	return serverTimeResponse.Time, nil
}

type MarketsResponse []MarketResponse

type MarketResponse struct {
	Id     Market
	Last   Amount
	High   Amount
	Change Amount
	Low    Amount
	Volume Amount
}

// GetMarkets gets all tradable market pairs.
func (c *Client) GetMarkets() (MarketsResponse, error) {
	var marketsResponse MarketsResponse
	url := BuildURL("/v2/markets")
	err := c.PublicGetJson(url, &marketsResponse)
	if err != nil {
		return marketsResponse, err
	}
	return marketsResponse, nil
}

// GetMarkets gets info for a specific market.
func (c *Client) GetMarket(market Market) (MarketResponse, error) {
	var marketResponse MarketResponse
	url := BuildURL("/v2/markets/%s", market)
	err := c.PublicGetJson(url, &marketResponse)
	if err != nil {
		return marketResponse, err
	}
	return marketResponse, nil
}

type TickerResponse struct {
	Market Market     `json:"market"`
	Bid    FiatAmount `json:"bid"`
	Ask    FiatAmount `json:"ask"`
	Spread FiatAmount `json:"spread"`
}

// GetTicket gets a market ticker for a specific market.
func (c *Client) GetTicker(market Market) (TickerResponse, error) {
	var tickerResponse TickerResponse
	url := BuildURL("/v2/markets/%s/ticker", market)
	err := c.PublicGetJson(url, &tickerResponse)
	if tickerResponse.Market == "" {
		tickerResponse.Market = market
	}
	return tickerResponse, err
}

type MarketHistory []MarketTrade

type MarketTrade struct {
	Type string `json:"type"`
	Amount Amount `json:"amount"`
	Price Amount `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	Total Amount `json:"total"`
}

type MilliTime time.Time

func (mt *MilliTime) Time() time.Time {
	return time.Time(*mt)
}

func (mt *MilliTime) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(time.RFC3339, s, time.UTC)
	*mt = MilliTime(t)
	return err
}

func (c *Client) GetMarketHistory(market Market, count int) (MarketHistory, error) {
	if count <= 0 || 1000 < count {
		return nil, errors.New("must have 0 < count <= 1000")
	}
	var marketHistory MarketHistory
	url := BuildURL("/v2/markets/%s/history?count=%d", market, count)
	err := c.PublicGetJson(url, &marketHistory)
	return marketHistory, err
}

type OrderBook struct {
	Bids []Order
	Asks []Order
}

type Order [2]Amount

func (o Order) Price() float64 {
	return float64(o[0])
}

func (o Order) Amount() float64 {
	return float64(o[1])
}

func (c *Client) GetMarketDepth(market Market) (OrderBook, error) {
	var orderBook OrderBook
	url := BuildURL("/v2/markets/%s/depth", market)
	err := c.PublicGetJson(url, &orderBook)
	return orderBook, err
}