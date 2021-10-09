package controller

import (
	"net/http"
	"os"
)

var TELEGRAM_BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")

func PostMessage(w http.ResponseWriter, r *http.Request) {
	/*var update, err = helpers.HandleRequest(r)

	if err != nil {
		log.Fatal(err)
		return
	}

	var _, errMess = sendMessage(update.Message.Chat.Id, update.Message.Text)
	if errMess != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("Message successfully delivered")
		return
	}*/
}
