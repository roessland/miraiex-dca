package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path"
)

func viperMustGetString(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Fatal(fmt.Sprintf("You must set config/env value %s", key))
	}
	return value
}

func getDbPath() string {
	configPath := viper.ConfigFileUsed()
	configDir := path.Dir(configPath)
	return path.Join(configDir, "miraiex-dca.db")
}
