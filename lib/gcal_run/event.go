package gcal_run

import (
	"fmt"
	"time"

	"google.golang.org/api/calendar/v3"
)

type Event struct {
	ID      string
	StartAt time.Time
	Summary string
	URL     string
}

func NewEvent(event *calendar.Event) (*Event, error) {
	isoFormatDataString := event.Start.DateTime
	// '2023-06-28T16:30:00+09:00' という形式のisoFormatDataString文字列を time.Datetime型に変換する
	startAt, err := time.Parse(time.RFC3339, isoFormatDataString)
	if err != nil {
		// 終日イベントなど、開始日時がない場合は、RFC3339の形式になっていない
		return nil, fmt.Errorf("faild to parse event.Start.Datetime = %s. %w", event.Start.DateTime, err)
	}
	url, err := NewURLParser().Parse(event)
	if url != "" {
		return &Event{
			ID:      event.Id,
			StartAt: startAt,
			Summary: event.Summary,
			URL:     url,
		}, nil
	}
	return nil, err
}

func (e *Event) TimeToStartSec() int {
	return int(e.StartAt.Sub(time.Now()).Seconds())
}

func (e *Event) String() string {
	return fmt.Sprintf("%s %s", e.StartAt.Format("2006-01-02 15:04"), e.Summary)
}
