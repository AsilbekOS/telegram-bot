package main

import (
	"bot/service/bot"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Botni yaratish va ishga tushirish
	b, err := bot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("bot started")
	b.Start()
}
