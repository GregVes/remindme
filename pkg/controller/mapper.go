package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	RequestMapper struct {
		Request *http.Request
		Update  *Update
	}
	Update struct {
		UpdateId int     `json:"update_id"`
		Message  Message `json:"message"`
	}
	Message struct {
		MessageId int    `json:"message_id"`
		Date      int    `json:"date"`
		Text      string `json:"text"`
		Chat      Chat   `json:"chat"`
		From      From   `json:"from"`
	}
	From struct {
		Id           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	}
	Chat struct {
		Id        int    `json:"id"`
		FirstName string `json:"first_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	}
)

func NewRequestMapper(r *http.Request) *RequestMapper {
	return &RequestMapper{
		Request: r,
	}
}

func (m *RequestMapper) MapToUpdate() error {
	var update Update
	if err := json.NewDecoder(m.Request.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming message %s", err.Error())
		return err
	}
	m.Update = &update
	return nil
}
