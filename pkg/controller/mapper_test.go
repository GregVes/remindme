package controller

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapToUpdate(t *testing.T) {

	var bodyStr = "{\"update_id\":29329212,\"message\":{\"message_id\":566,\"from\":{\"id\":1603037541,\"is_bot\":false,\"first_name\":\"Greg\",\"username\":\"gregentoo\",\"language_code\":\"e\"},\"chat\":{\"id\":1633037542,\"first_name\":\"Greg\",\"username\":\"gregentoo\",\"type\":\"private\"},\"date\":1633703801,\"text\":\"hello world\"}}"
	readCloserBody := ioutil.NopCloser(bytes.NewBuffer([]byte(bodyStr)))

	tests := []struct {
		input *http.Request
		want  *Update
	}{
		{
			input: &http.Request{
				Body: readCloserBody,
			},
			want: &Update{
				UpdateId: 29329212,
				Message: Message{
					MessageId: 566,
					Date:      1633703801,
					Text:      "hello world",
					Chat: Chat{
						Id:        1633037542,
						FirstName: "Greg",
						Username:  "gregentoo",
						Type:      "private",
					},
					From: From{
						Id:           1603037541,
						IsBot:        false,
						FirstName:    "Greg",
						Username:     "gregentoo",
						LanguageCode: "e",
					},
				},
			},
		},
	}
	for _, tc := range tests {
		mapper := NewRequestMapper(tc.input)
		got := mapper.MapToUpdate()
		if got != nil {
			t.Fatal(got)
		}
		assert.Equal(t, tc.want, mapper.Update)
	}
}
