package model

// body of Bot API request
type (
	Update struct {
		UpdateId int `json:"update_id"`
		Message Message `json:"message"`
	}
	Message struct {
		MessageId int `json:"message_id"`
		Date int `json:"date"`
		Text string `json:"text"`
		Chat Chat `json:"chat"`
		From From `json:"from"`
	}
	From struct {
		Id int `json:"id"`
		IsBot bool `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username string `json:"username"`
		LanguageCode string `json:"language_code"`
	}
	Chat struct {
		Id int `json:"id"`
		FirstName string `json:"first_name"`
		Username string `json:"username"`
		Type string `json:"type"`
	}
)