package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNewReminder(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "/remindme check the stock price | October 17 @ 2:30AM",
			want:  true,
		},
		{
			input: "/reminyme check the stock price | October 17 @ 1:10AM",
			want:  false,
		},
		{
			input: "check the stock price | October 17 @ 5pm /remindme",
			want:  false,
		},
	}
	for _, tc := range tests {
		got := IsNewReminder(tc.input)
		assert.Equal(t, tc.want, got)
	}
}

func TestIsValidInput(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "/remindme check the stock price | October 17 @ 1:01AM",
			want:  true,
		},
		{
			input: "/remindme check the stock price October 17 @ 1:01AM",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 17 1:01AM",
			want:  false,
		},
		{
			input: "/remindme October 18 | check the stock price @ 1:01AM",
			want:  false,
		},
		{
			input: "/remindme check the stock price October | 17 @ 1:01AM",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 17 1:01AM",
			want:  false,
		},
		{
			input: "/remindme check the stock price | 18 October @ 1:01AM",
			want:  false,
		},
		{
			input: "/remindme @ 15:00 check the stock price",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 26",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 26 @ 1:01AM",
			want:  true,
		},
		{
			input: "/remindme check the stock price | October 26 @ 23:01",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 26 @ 443:434",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 26 @ 484828",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 26 @ 484828",
			want:  false,
		},
	}
	for _, tc := range tests {
		converter := NewConverter(tc.input)
		got := converter.IsValidInput()
		assert.Equal(t, tc.want, got)
	}
}

func TestPatternSearch(t *testing.T) {
	tests := []struct {
		pattern string
		input   string
		want    bool
	}{
		{
			pattern: "(hello)",
			input:   "hello",
			want:    true,
		},
		{
			pattern: "(each\\s+sunday)",
			input:   "each sunday",
			want:    true,
		},
		{
			pattern: "(each\\s+sunday)",
			input:   "eachsunday",
			want:    false,
		},
		{
			pattern: "October 26",
			input:   "26 October",
			want:    false,
		},
		{
			pattern: "(tomorrow)|each\\s+sunday",
			input:   "tomorrow",
			want:    true,
		},
	}
	for _, tc := range tests {
		got, _ := patternSearch(tc.pattern, tc.input)
		assert.Equal(t, tc.want, got)
	}
}

func TestGetRawReminder(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "/remindme check the stock price | October 17 @ 1:01AM",
			want:  "check the stock price | October 17 @ 1:01AM",
		},
		{
			input: "/remindme     buy bread | tomorrow @ 6:09PM",
			want:  "buy bread | tomorrow @ 6:09PM",
		},
	}
	for _, tc := range tests {
		converter := NewConverter(tc.input)
		converter.ExtractRawReminder()
		assert.Equal(t, tc.want, converter.Input)
	}
}

/*func TestToReminder(t *testing.T) {

	targetDate, err := time.Parse(constants.FullDateFormat, "October 17 2021")

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		input string
		want  *repo.Reminder
	}{
		{
			input: "/remindme check the stock price : October 17 @ 1:01AM",
			want: &repo.Reminder{
				Id:           uuid.New().String(),
				ChatId:       1111111,
				ChatMessage:  "check the stock price",
				RecurrentDay: "Tuesday",
				TargetDate:   targetDate,
				TargetTime:   "17:00",
			},
		},
	}
	for _, tc := range tests {
		converter := NewConverter(tc.input)
		got := converter.ToReminder()
		if got != nil {
			t.Fatal(got)
		}
		assert.Equal(t, tc.want, converter.Reminder)
	}
}*/
