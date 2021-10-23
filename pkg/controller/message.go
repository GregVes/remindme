package controller

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gregves/remindme/pkg/constants"
	postgresql "github.com/gregves/remindme/pkg/postgresql"
	converter "github.com/gregves/remindme/pkg/service"
)

var SendMesssageEndpoint = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"

type (
	Reminder struct {
	}
	storage interface {
		SaveReminder(*Reminder)
	}
)

func PostMessage(w http.ResponseWriter, r *http.Request) {
	mapper := NewRequestMapper(r)
	err := mapper.MapToUpdate()
	if err != nil {
		log.Print(err)
		return
	}

	chatId := mapper.Update.Message.Chat.Id
	chatMessage := mapper.Update.Message.Text

	if !converter.IsNewReminder(chatMessage) {
		Request(chatId, constants.ErrNotAReminder.Error())
		return
	}

	converter := converter.NewConverter(chatMessage)

	converter.ExtractRawReminder()

	err = converter.IsValidInput()
	if err != nil {
		Request(chatId, err.Error())
		return
	}

	converter.ToReminder(chatId)
	log.Print(os.Getenv("REMINDME_DB_DSN"))
	repo, err := postgresql.NewRepository("postgres", os.Getenv("REMINDME_DB_DSN"), 2, 2)

	if err != nil {
		log.Print(err)
		Request(chatId, constants.ErrDb.Error())
		return
	}

	err = repo.Save(&converter.Reminder)

	if err != nil {
		log.Print(err)
		Request(chatId, constants.ErrDb.Error())
		return
	}
	Request(chatId, constants.SuccessSave)

	defer repo.Close()
}

func Request(chatId int, message string) bool {
	data := url.Values{
		"chat_id": {strconv.Itoa(chatId)},
		"text":    {message},
	}

	_, err := http.PostForm(SendMesssageEndpoint, data)

	if err != nil {
		log.Printf("error when posting text to the chat with id %d: %s", chatId, err.Error())
		return false
	}
	log.Printf("Message '%s' successfully sent to chat %d", message, chatId)
	return true
}
