package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gregves/remindme/telegram-bot-api"
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

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert(WEBHOOK_ENDPOINT, SSL_CERT))

	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()

	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		update, err := bot.HandleUpdate(r)
		if err != nil {
			log.Printf("%+v\n", err.Error())
		} else {
			log.Printf("%+v\n", *update)
		}
	})

	http.ListenAndServe(":8002", nil)

	for update := range updates {
		log.Println("%+v\n", update)
	}
}
