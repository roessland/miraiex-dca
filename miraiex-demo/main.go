package main

import (
	"fmt"
	"github.com/roessland/miraiex-dca/miraiex"
	"log"
	"time"
)

func main() {
	// Create a new MiraiEx client
	client := miraiex.NewClient()
	// Configure API access keys.
	// Get them here: https://platform.miraiex.com/client/settings/apikey
	// Enable Write permission at your own peril.
	client.ApiKey = ApiKey
	client.ClientID = ClientID
	client.SecretKey = SecretKey

	// Get current server time
	millis, err := miraiex.GetServerTime()
	fmt.Println("\nCurrent server time is ", millis)

	// Get account balances for all currencies
	balances, err := client.GetBalances()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nAccount balances:")
	fmt.Println("\tName\tAvailable\tBalance\t          Hold")
	for _, balance := range balances {
		fmt.Printf("\t%s\t%8.4f\t%8.4f\t%8.4f\n",
			balance.Currency, balance.Available, balance.Balance, balance.Hold)
	}

	// Get all tradable markets
	markets, err := miraiex.GetMarkets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nAll tradable markets")
	for _, market := range markets {
		fmt.Printf("\t%#v\n", market)
	}

	// Get a specific market
	market, err := miraiex.GetMarket(miraiex.BTCNOK)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nOne specific market")
	fmt.Printf("\t%#v\n", market)

	// Get current price of BTC
	btcTicker, err := miraiex.GetTicker(miraiex.BTCNOK)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nCurrent BTC price")
	fmt.Printf("\tBid: %.2f,  Ask: %.2f,  Spread: %.2f\n", btcTicker.Bid, btcTicker.Ask, btcTicker.Spread)

	// Get trade history for BTCNOK market
	marketHistory, err := miraiex.GetMarketHistory(miraiex.BTCNOK, 5)
	if err != nil {
		log.Fatal(err)
	}
	for _, trade := range marketHistory {
		fmt.Printf("\t%#v\t%s\n", trade, trade.CreatedAt.Format(time.RFC3339Nano))
	}

	// Get order book for BTCNOK market
	orderBook, err := miraiex.GetMarketDepth(miraiex.BTCNOK)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Asks:\n", orderBook.Asks)
	fmt.Println("Bids:\n", orderBook.Bids)

	// Buy BTC. Uncomment the following lines to buy the minimum allowed amount of Bitcoin.
	//btcPrice := btcTicker.Ask
	//var btcAmount miraiex.CryptoAmount = 0.0001
	//orderId, err := client.CreateOrder(miraiex.BTCNOK, miraiex.BidMaxPrice, btcPrice, btcAmount)
	//totalPrice := miraiex.FiatAmount(btcAmount) * btcPrice
	//log.Printf("Created order %s for %.6f mBTC @ %.2f NOK, for a total price of %.2f NOK + 0.5%% fees", orderId, 1000*btcAmount, btcPrice, totalPrice)
}
