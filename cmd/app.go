package main

import (
	"time"

	"github.com/VicSobDev/sigamonitor/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

// Define a struct named BotApp to hold application data and configuration.
type BotApp struct {
	messageMap map[int64]bot.District // Store message data
	config     Config                 // Store configuration data
	logger     *zap.Logger            // Logger for logging messages
	tgBot      *tgbotapi.BotAPI       // Telegram Bot API client
}

// NewBotApp initializes and returns a new BotApp instance.
func NewBotApp() (*BotApp, error) {
	// Create a new BotApp instance with initializations
	app := &BotApp{
		messageMap: make(map[int64]bot.District),
	}

	// Load configuration settings
	var err error
	app.config, err = loadConfig()
	if err != nil {
		return nil, err
	}

	// Initialize a production logger
	app.logger, err = zap.NewProduction()
	if err != nil {
		return nil, err
	}

	// Initialize the Telegram Bot API client
	app.tgBot, err = tgbotapi.NewBotAPI(app.config.BotToken)
	if err != nil {
		return nil, err
	}
	app.logger.Info("Authorized on account", zap.String("username", app.tgBot.Self.UserName))

	return app, nil
}

// Run starts the main loop of the application.
func (app *BotApp) Run() {
	app.sendBotStartMessage() // Send a start message

	// Create a ticker to execute the processCheck function periodically
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	// Start the main loop
	for range ticker.C {
		//Start the monitor
		app.processCheck()
	}
}
