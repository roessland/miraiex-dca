package miraiex

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Market is a market available on MiraiEx
type Market string

const BTCNOK Market = "BTCNOK"
const LTCNOK Market = "LTCNOK"

// OrderType is either a bid or an ask.
type OrderType string

const BidMaxPrice OrderType = "Bid"
const AskMinPrice OrderType = "Ask"

func OrderTypeFromString(name string) OrderType {
	switch strings.ToLower(name) {
	case "bid":
		return BidMaxPrice
	case "ask":
		return AskMinPrice
	}
	return ""
}

// Amount is either fiat or crypto amount (based on context)
type Amount float64

// FiatAmount is an amount of NOK, USD, etc.
type FiatAmount float64

// CryptoAmount is an amount of BTC or LTC.
type CryptoAmount float64

func (am *Amount) UnmarshalJSON(data []byte) error {
	floatVal, err := unmarshalStringToFloat(data)
	*am = Amount(floatVal)
	return err
}

func (am *FiatAmount) UnmarshalJSON(data []byte) error {
	floatVal, err := unmarshalStringToFloat(data)
	*am = FiatAmount(floatVal)
	return err
}

func (am *CryptoAmount) UnmarshalJSON(data []byte) error {
	floatVal, err := unmarshalStringToFloat(data)
	*am = CryptoAmount(floatVal)
	return err
}

func unmarshalStringToFloat(data []byte) (float64, error) {
	var err error
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return -1.0, err
	}

	floatAmount, err := strconv.ParseFloat(v, 10)
	return floatAmount, err
}