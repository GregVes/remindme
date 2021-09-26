package main

import (
	"os"
	"log"
	"net/http"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var TELEGRAM_BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")
var WEBHOOK_ENDPOINT = "https://radio4000-dev-api.space/bot"
var SSL_CERT = "fullchain.pem"

func main() {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_BOT_TOKEN)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Print("Authorized on account %s", bot.Self.UserName)

	_, err = bot.Request(tgbotapi.NewWebhookWithCert(WEBHOOK_ENDPOINT, SSL_CERT))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	info, err := bot.GetWebhookInfo()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}

	log.Println(info)

	_ = bot.ListenForWebhook("/")

	http.ListenAndServe(":8002", nil)

	/*for update := range updates {
		log.Println("%+v\n", update)
	}*/
}
