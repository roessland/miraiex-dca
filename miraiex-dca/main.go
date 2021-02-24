package main

import (
	"fmt"
	"github.com/roessland/miraiex-dca/miraiex-dca/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"

	"github.com/roessland/miraiex-dca/miraiex"
	"github.com/roessland/miraiex-dca/miraiex-dca/storage"
)

/*
TODO: CLI parsing.
TODO: Config parsing.
TODO: Database. (Bolt?)
*/

const apiKeyConfigKey = "MIRAIEX_API_KEY"
const clientIDConfigKey = "MIRAIEX_CLIENT_ID"
const secretKeyConfigKey = "MIRAIEX_SECRET_KEY"

func main() {
	// Set config file search locations
	viper.SetConfigFile("miraiex-dca.yaml")
	viper.AddConfigPath("/etc/miraiex-dca/")
	viper.AddConfigPath("$HOME/.miraiex-dca/")
	viper.AddConfigPath(".")

	// Allow overriding values with environment variables
	viper.SetDefault(apiKeyConfigKey, "")
	viper.SetDefault(clientIDConfigKey, "")
	viper.SetDefault(secretKeyConfigKey, "")
	viper.AutomaticEnv()

	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error loading config file: %s \n", err))
	}

	// Initiate logging
	log := logrus.New()

	// Initiate database
	dbPath := getDbPath()
	repo, err := storage.NewRepo(dbPath)
	if err != nil {
		panic(fmt.Errorf("Fatal error opening database from disk: %s \n", err))
	}
	defer repo.Close()

	// Initiate MiraiEx API client
	apiKey := viperMustGetString(apiKeyConfigKey)
	clientID := viperMustGetString(clientIDConfigKey)
	secretKey := viperMustGetString(secretKeyConfigKey)
	mxClient := miraiex.NewClient(log).SetAuthentication(apiKey, clientID, secretKey)
	balances, err := mxClient.GetBalances()
	log.WithField("balances", balances).Info()

	// Get current price of BTC
	btcTicker, err := mxClient.GetTicker(miraiex.BTCNOK)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nCurrent BTC price")
	fmt.Printf("\tBid: %.2f,  Ask: %.2f,  Spread: %.2f\n", btcTicker.Bid, btcTicker.Ask, btcTicker.Spread)

	// Buy BTC. Uncomment the following lines to buy the minimum allowed amount of Bitcoin.
	btcPrice := btcTicker.Ask
	var btcAmount miraiex.CryptoAmount = 0.0001
	totalPrice := miraiex.FiatAmount(btcAmount) * btcPrice
	logFields := log.WithFields(logrus.Fields{
		"market":        miraiex.BTCNOK,
		"order_type":    miraiex.BidMaxPrice,
		"price_per_btc": btcPrice,
		"amount_btc":    btcAmount,
		"total_price":   totalPrice,
	})
	logFields.WithField("event", "miraiex.order.create.start").Info()
	orderId, err := mxClient.CreateOrder(miraiex.BTCNOK, miraiex.BidMaxPrice, btcPrice, btcAmount)
	logFields = logFields.WithField("order_id", orderId)
	success := orderId != "" && err == nil
	if !success {
		logFields.WithField("event", "miraiex.order.create.failed").WithError(err).Error()
	} else {
		logFields.WithField("event", "miraiex.order.create.success").Info()
	}

	// Save order to database
	order := models.Order{
		ExternalID: orderId,
		OrderType:  string(miraiex.BidMaxPrice),
		Price:      float64(btcPrice),
		Amount:     float64(btcAmount),
		Market:     string(miraiex.BTCNOK),
		Time:       time.Now(),
	}
	_, err = repo.CreateOrder(&order)
	if err != nil {
		panic(err)
	}

	// Parse flags
	// Read config + secrets.
	// For each missing config option, ask for user input.
	// Save config to home directory.
	// Init database
	// Do migrations

	// switch on subcommand
}

func cron() {
	// check if should buy bitcoin according to config, and how much
	// check status of WIP orders
	// check account balance in NOK
	// buy bitcoin
}

func web() {
	// host webserver
}
