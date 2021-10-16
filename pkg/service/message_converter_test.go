package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfMessageIsNewReminder(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "/remindme check the stock price | October 17 @ 5pm",
			want:  true,
		},
		{
			input: "/reminyme check the stock price | October 17 @ 5pm",
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

/*func TestConvertMessageStringIntoReminderObject(t *testing.T) {

	targetDate, err := time.Parse(constants.FullDateFormat, "October 17 2021")

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		input string
		want  *repo.Reminder
	}{
		{
			input: "/remindme check the stock price : October 17 @ 5pm",
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

func TestValidateInputMessage(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		/*{
			input: "/remindme check the stock price | October 17 @ 15:00",
			want:  true,
		},
		{
			input: "/remindme check the stock price October 17 @ 15:00",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 17 15:0ß",
			want:  false,
		},
		{
			input: "/remindme October 18 | check the stock price @ 15:00",
			want:  false,
		},
		{
			input: "/remindme check the stock price October | 17 @ 15:00",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 17 15:0ß",
			want:  false,
		},
		{
			input: "/remindme check the stock price | 18 October @ 15:00",
			want:  false,
		},
		{
			input: "/remindme @ 15:00 check the stock price",
			want:  false,
		},
		{
			input: "/remindme check the stock price | October 26",
			want:  false,
		},*/
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
			input: "/remindme check the stock price | October 26 @ super time",
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
		got, _ := PatternSearch(tc.pattern, tc.input)
		assert.Equal(t, tc.want, got)
	}
}
