package main

import (
	//"log"

	"fmt"
	"net/http"
	"os"

	"database/sql"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/gregves/remindme/pkg/controller"
	_ "github.com/lib/pq"
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
	migrateDb("postgres", os.Getenv("REMINDME_DB_DSN"), os.Getenv("DB_MIGRATIONS_DIR"))

	r := mux.NewRouter()

	r.HandleFunc("/bot", controller.PostMessage).Methods("POST")
	r.HandleFunc("/batch", controller.PostReminders).Methods("POST")
	http.Handle("/bot", r)
	http.Handle("/batch", r)

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

func migrateDb(dbType string, dsn string, migrationsDir string) {
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsDir, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	m.Steps(2)
	log.Info("Migration - if any - done")
}
