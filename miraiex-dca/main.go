package main

import (
	"fmt"
	"github.com/roessland/miraiex-dca/miraiex-dca/models"
	"log"

	"github.com/spf13/viper"

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
	mxClient := miraiex.NewClient().SetAuthentication(apiKey, clientID, secretKey)
	fmt.Println(mxClient.GetBalances())

	// Print all orders in DB
	orders, err := repo.GetOrders()
	if err != nil {
		panic(err)
	}
	for _, o := range orders {
		fmt.Println("DB order: ", o)
	}

	// Get current price of BTC
	btcTicker, err := miraiex.GetTicker(miraiex.BTCNOK)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nCurrent BTC price")
	fmt.Printf("\tBid: %.2f,  Ask: %.2f,  Spread: %.2f\n", btcTicker.Bid, btcTicker.Ask, btcTicker.Spread)

	// Buy BTC. Uncomment the following lines to buy the minimum allowed amount of Bitcoin.
	btcPrice := btcTicker.Ask
	var btcAmount miraiex.CryptoAmount = 0.0001
	orderId, err := mxClient.CreateOrder(miraiex.BTCNOK, miraiex.BidMaxPrice, btcPrice, btcAmount)
	totalPrice := miraiex.FiatAmount(btcAmount) * btcPrice
	log.Printf("Created order %s for %.6f mBTC @ %.2f NOK, for a total price of %.2f NOK + 0.5%% fees", orderId, 1000*btcAmount, btcPrice, totalPrice)

	// Save order to database
	order := models.Order{
		ExternalID: orderId,
		OrderType:  string(miraiex.BidMaxPrice),
		Price:      float64(btcPrice),
		Amount:     float64(btcAmount),
		Market:     string(miraiex.BTCNOK),
	}
	err = repo.CreateOrder(&order)
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
