package miraiex

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) PrivateGetJson(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("miraiex-access-key", c.ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Print("Response body:", string(body))

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		return errors.New(fmt.Sprintf("Got status code %d", resp.StatusCode))
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) PrivatePostJson(url string, requestV, responseV interface{}) error {
	reqBody, err := json.Marshal(requestV)
	if err != nil {
		return err
	}
	log.Print("Request body: ", string(reqBody))

	req, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Add("miraiex-access-key", c.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do: %w", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Print("Response body:", string(body))

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		return errors.New(fmt.Sprintf("Got status code %d", resp.StatusCode))
	}

	err = json.Unmarshal(body, responseV)
	if err != nil {
		return err
	}
	return nil
}

type BalancesResponse []Balance

type Balance struct {
	Currency  string  `json:"currency"`
	Balance   Amount `json:"balance"`
	Hold      Amount `json:"hold"`
	Available Amount `json:"available"`
}

func (c *Client) GetBalances() (BalancesResponse, error) {
	var balancesResponse BalancesResponse
	url := BuildURL("/v2/balances")
	err := c.PrivateGetJson(url, &balancesResponse)
	if err != nil {
		return balancesResponse, err
	}
	return balancesResponse, nil
}

type CreateOrderRequest struct {
	Market string `json:"market"`
	Type   string `json:"type"`
	Price  string `json:"price"`
	Amount string `json:"amount"`
}

type CreateOrderResponse struct {
	Id json.Number `json:"id"`
}

// CreateOrder buys or sells.
// Response code 400 can mean: Invalid decimal formatting, too high price, not enough available balance in that market.
func (c *Client) CreateOrder(market Market, orderType OrderType, priceFiat FiatAmount, amountCrypto CryptoAmount) (string, error) {
	if priceFiat < 0.01 {
		return "", fmt.Errorf("lowest allowed priceFiat is 0.01, but got %f", priceFiat)
	}
	if amountCrypto < 0.0001 {
		return "", fmt.Errorf("lowest allowed amountCrypto is 0.0001, but got %f", amountCrypto)
	}
	req := CreateOrderRequest{}
	req.Market = string(market)
	req.Type = string(orderType)
	req.Price = fmt.Sprintf("%.2f", priceFiat)
	req.Amount = fmt.Sprintf("%f", amountCrypto)

	resp := CreateOrderResponse{}
	url := BuildURL("/v2/orders")
	err := c.PrivatePostJson(url, req, &resp)
	return string(resp.Id), err
}
