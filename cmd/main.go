package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/gorilla/mux"
	"github.com/gregves/remindme/pkg/controller"
)

var PORT = 8002

func main() {

	initGraylog()

	r := mux.NewRouter()

	r.HandleFunc("/bot", controller.PostMessage).Methods("POST")
	http.Handle("/bot", r)

	log.Println(fmt.Sprintf("Starting server at port %d", PORT))

	log.Print(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))

	// parse db every 30mn to check for reminders to push to chats. Maybe it will be a separate application that the API should talk to
	// if matches, push reminders to charts
}

func initGraylog() {
	graylogAddr := os.Getenv("GRAYLOG_ENDPOINT")

	if graylogAddr != "" {
		gelfWriter, err := gelf.NewWriter(graylogAddr)
		if err != nil {
			log.Printf("gelf.NewWriter: %s", err)
		} else {
			log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
			log.Printf("Starting to log to '%s'", graylogAddr)
		}
	} else {
		log.Print("Missing GRAYLOG_ENDPOINT env var")
	}
}
