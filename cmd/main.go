package main

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/gregves/remindme/pkg/controller"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.PostMessage).Methods("POST")
	http.Handle("/", r)
}
