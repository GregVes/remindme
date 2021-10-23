package controller

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

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
	// map request to Update

	// get text
	// parse text into a Reminder object (type + message + occurence + time)
	// store reminder into db

	mapper := NewRequestMapper(r)
	err := mapper.MapToUpdate()
	if err != nil {
		log.Print(err)
		return
	}

	chatId := mapper.Update.Message.Chat.Id
	chatMessage := mapper.Update.Message.Text

	if !converter.IsNewReminder(chatMessage) {
		Request(chatId, "Not a reminder. Prefix your message with /remindme")
		return
	}

	converter := converter.NewConverter(chatMessage)

	converter.ExtractRawReminder()

	err = converter.IsValidInput()
	if err != nil {
		if !Request(chatId, err.Error()) {
			return
		}
	}

	//defer response.Body.Close()

	// err = converter.ToReminder(chatId)
	// if err != nil {
	// 	log.Print(err)
	// }
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", "remindme", os.Getenv("REMINDME_DB_PASSWORD"), os.Getenv("REMINDME_DB_HOST"), "remindme")
	// repository, err := postgresql.NewRepository("postgres", dsn, 2, 2)

	// if err != nil {
	// 	log.Print(err)
	// }
	// err = repository.Save(&converter.Reminder)

	// defer repository.Close()

	// if err != nil {
	// 	log.Print(err)
	// }
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
	log.Printf("Message successfully sent to chat with id %d", chatId)
	return true
}
