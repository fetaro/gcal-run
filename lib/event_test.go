package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/calendar/v3"
)

func TestNewEvent(t *testing.T) {
	event := calendar.Event{
		ConferenceData: &calendar.ConferenceData{
			EntryPoints: []*calendar.EntryPoint{
				{
					EntryPointType: "video",
					Uri:            "https://meet.google.com/xxx",
				},
			},
		},
		Id:      "123",
		Summary: "test",
		Start: &calendar.EventDateTime{
			DateTime: "2023-12-28T22:05:00+09:00",
		},
	}

	// テスト対象の関数を実行
	actual, err := NewEvent(&event)

	// テスト結果の判定
	assert.NoError(t, err)
	assert.Equal(t, actual.ID, "123")
	assert.Equal(t, actual.Summary, "test")
	assert.Equal(t, actual.URL, "https://meet.google.com/xxx")
	assert.True(t, actual.TimeToStartSec() < 0)
	assert.Equal(t, actual.MeetingApp, "Meet")
	assert.Equal(t, fmt.Sprintf("%s", actual), "2023-12-28 22:05 test")
}

func TestNewEventFail(t *testing.T) {
	event := calendar.Event{
		Id:      "123",
		Summary: "test",
		Start: &calendar.EventDateTime{
			DateTime: "InvalidDataFormat",
		},
	}

	// テスト対象の関数を実行
	actual, err := NewEvent(&event)

	// テスト結果の判定
	assert.Error(t, err)
	assert.Nil(t, actual)
}
