package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(GetServerTime())
	client := NewClient()
	client.ApiKey = ApiKey
	client.ClientID = ClientID
	client.SecretKey = SecretKey

	balances, _ := client.GetBalances()
	fmt.Printf("%#v\n", balances)

	// Get current price
	btcTicker, err := GetTicker(BTCNOK)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Current BTC price: ", btcTicker)
	btcPrice := btcTicker.Ask

	// Buy BTC
	var btcAmount CryptoAmount = 0.0001
	orderId, err := client.CreateOrder(BTCNOK, BidMaxPrice, btcPrice, btcAmount)
	totalPrice := FiatAmount(btcAmount) * btcPrice
	log.Printf("Created order %s for %.6f mBTC @ %.2f NOK, for a total price of %.2f NOK + 0.5%% fees", orderId, 1000*btcAmount, btcPrice, totalPrice)
}
