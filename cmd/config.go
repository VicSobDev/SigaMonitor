package main

import (
	"fmt"
	"os"
	"strconv"
)

// Define a struct named Config to hold configuration settings.
type Config struct {
	ChannelID      string
	ChannelIDAsInt int
	BotToken       string
}

// loadConfig reads environment variables and populates the Config struct.
func loadConfig() (Config, error) {
	var config Config

	// Read the CHANNEL_ID environment variable
	config.ChannelID = os.Getenv("CHANNEL_ID")

	// Convert ChannelID to an integer
	var err error
	config.ChannelIDAsInt, err = strconv.Atoi(config.ChannelID)
	if err != nil {
		return config, fmt.Errorf("can't convert channel id to int: %v", err)
	}

	// Read the BOT_TOKEN environment variable
	config.BotToken = os.Getenv("BOT_TOKEN")

	return config, nil
}
