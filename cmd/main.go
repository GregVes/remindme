package main

import (
	//"log"

	"fmt"
	"net/http"
	"os"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/gorilla/mux"
	"github.com/gregves/remindme/pkg/controller"
	log "github.com/sirupsen/logrus"
)

type NullFormatter struct {
}

// Don't spend time formatting logs
func (NullFormatter) Format(e *log.Entry) ([]byte, error) {
	return []byte{}, nil
}

var PORT = 8002

func main() {
	initGraylog()

	r := mux.NewRouter()

	r.HandleFunc("/bot", controller.PostMessage).Methods("POST")
	http.Handle("/bot", r)

	log.Info(fmt.Sprintf("Starting server at port %d", PORT))

	log.Info(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func initGraylog() {
	graylogAddr := os.Getenv("GRAYLOG_ENDPOINT")
	hook := graylog.NewGraylogHook(graylogAddr, map[string]interface{}{})
	log.SetReportCaller(true)
	log.AddHook(hook)
	log.Info("Starting to send application logs to Graylog instance")
	log.SetFormatter(new(NullFormatter)) // Don't send logs to stdout
}
