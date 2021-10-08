package util

import(
  "log"
	"net/http"
	"encoding/json"
	"github.com/gregves/remindme/pkg/model"
)

func parseRequest(r *http.Request) (*model.Update, error) {
	var update model.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming message %s", err.Error())
		return nil, err
	}
	log.Println(r.Body)
	return &update, nil
}