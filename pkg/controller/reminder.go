package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	repo "github.com/gregves/remindme/pkg/repository"
	log "github.com/sirupsen/logrus"
)

//var SendMesssageEndpoint = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"

type ()

func PostReminders(w http.ResponseWriter, r *http.Request) {
	// resquest body = []Reminders
	// iterate over reminders
	// /sendMessage to reminder.ChatId with reminder.ChatMessage as reponse.Body
	var reminders repo.Reminders
	var jsonData []byte
	w.Header().Set("Content-Type", "application/json")

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Print("Cannot unmarshal request body into Reminders struct")
		w.WriteHeader(http.StatusBadRequest)
		jsonData = []byte(`{"status": "Cannot unmarshal request body into Reminders struct"}`)
		w.Write(jsonData)
		return
	}

	json.Unmarshal(data, &reminders)

	for _, reminder := range reminders.Data {
		isDone := Request(reminder.ChatId, reminder.ChatMessage)
		// you should thing of what happens if a reminder could not be delivered
		if !isDone {
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	jsonData = []byte(`{"status": "Reminders sent to users"}`)
	w.Write(jsonData)
}
