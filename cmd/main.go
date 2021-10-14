package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/gregves/remindme/pkg/controller"
)

var PORT = 8003

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/bot", controller.PostMessage).Methods("POST")
	http.Handle("/bot", r)

	log.Println(fmt.Sprintf("Starting server at port %d", PORT))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))

	// parse db every 30mn to check for reminders to push to chats. Maybe it will be a separate application that the API should talk to
	// if matches, push reminders to charts
}
