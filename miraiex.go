package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const ApiBase = "https://api.miraiex.com"

type Market string

const BTCNOK Market = "BTCNOK"
const LTCNOK Market = "LTCNOK"

type OrderType string

const BidMaxPrice OrderType = "Bid"
const AskMinPrice OrderType = "Ask"

type FiatAmount float64

type CryptoAmount float64

func BuildURL(path string, args ...interface{}) string {
	if args == nil {
		return ApiBase + path
	}
	return ApiBase + fmt.Sprintf(path, args...)
}

type Client struct {
	ApiKey    string
	ClientID  string
	SecretKey string
}

func NewClient() *Client {
	return &Client{}
}


func (fiat *FiatAmount) UnmarshalJSON(data []byte) error {
	var err error
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	floatAmount, err := strconv.ParseFloat(v, 10)
	*fiat = FiatAmount(floatAmount)
	if err != nil {
		return err
	}

	return nil
}
