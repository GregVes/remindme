package main

import (
	"os"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

var TELEGRAM_BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")

type (
	Update struct {
		UpdateId int `json:"update_id"`
		Message Message `json:"message"`
	}
	Message struct {
		Text string `json:"text"`
		Chat Chat `json:"chat"`
	}
	Chat struct {
		Id int `json:"id"`
	}
)

func handleRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming message %s", err.Error())
		return nil, err
	}
	log.Println(r.Body)
	return &update, nil
}

func HandleHook(w http.ResponseWriter, r *http.Request) {
	log.Print("received request")
	var update, err = handleRequest(r)

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
	}
}

func sendMessage(chatId int, text string) (string, error) {
	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + TELEGRAM_BOT_TOKEN + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}


func main() {
	/*var telegramApi string = "https://api.telegram.org/bot" + TELEGRAM_BOT_TOKEN + "/setWebhook?url=https://gregentoo.com/bot"
	response, err := http.Get(telegramApi)

	if err != nil {
		log.Printf("error when trying to set webhook: %s", err.Error())
		os.Exit(1)
	}
	defer response.Body.Close*/
	http.HandleFunc("/", HandleHook)
	log.Println("Starting server at port 8002")
	log.Fatal(http.ListenAndServe(":8002", nil))
}
