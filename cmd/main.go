package main

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Configure log format
}

func main() {
	app, err := NewBotApp()
	if err != nil {
		log.Fatalf("Failed to initialize the bot app: %v", err)
	}
	app.Run()
}
