package main

import (
	"fmt"
	"log"
	"os"

	postgresql "github.com/gregves/remindme/pkg/postgresql"
	converter "github.com/gregves/remindme/pkg/service"
)

var PORT = 8003

func main() {
	/*r := mux.NewRouter()

	r.HandleFunc("/bot", controller.PostMessage).Methods("POST")
	http.Handle("/bot", r)

	log.Println(fmt.Sprintf("Starting server at port %d", PORT))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))*/

	// parse db every 30mn to check for reminders to push to chats. Maybe it will be a separate application that the API should talk to
	// if matches, push reminders to charts

	input := "/remindme go to the counselor | each October 20 @ 23:00"
	chatId := 111
	if !converter.IsNewReminder(input) {
		log.Fatal("Not a reminder. Prefix your message with /remindme")
		os.Exit(1)
	}
	converter := converter.NewConverter(input)
	converter.ExtractRawReminder()
	isValid := converter.IsValidInput()
	if !isValid {
		log.Fatal("Invalid input")
		os.Exit(1)
	}

	err := converter.ToReminder(chatId)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5433/%s?sslmode=disable", "remindme", os.Getenv("REMINDME_DB_PASSWORD"), os.Getenv("REMINDME_DB_HOST"), "remindme")
	log.Print(dsn)
	repository, err := postgresql.NewRepository("postgres", dsn, 2, 2)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = repository.Save(&converter.Reminder)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
