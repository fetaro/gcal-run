package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/calendar/v3"
)

// ParseGoogleMeetEventUrlのテスト
func TestParseGoogleMeetEventUrl(t *testing.T) {
	event := &calendar.Event{
		ConferenceData: &calendar.ConferenceData{
			EntryPoints: []*calendar.EntryPoint{
				{
					EntryPointType: "video",
					Uri:            "https://meet.google.com/xxx",
				},
			},
		},
	}
	meetingApp, url := NewURLParser().Parse(event)

	assert.Equal(t, meetingApp, "Meet")
	assert.Equal(t, url, "https://meet.google.com/xxx")
}

func TestParseZoomEventUrlWhenLocationHasUrl(t *testing.T) {
	event := &calendar.Event{
		Location: "https://zoom.us/j/xxx",
	}
	meetingApp, url := NewURLParser().Parse(event)

	assert.Equal(t, meetingApp, "Zoom")
	assert.Equal(t, url, "https://zoom.us/j/xxx")
}

func TestParseZoomEventUrlWhenDescriptionHasUrl(t *testing.T) {
	event := &calendar.Event{
		Description: "<a href=\"https://zoom.us/j/12345?pwd=xxx\"><u>https://zoom.us/j/12345?pwd=xxx</u></a>",
	}
	meetingApp, url := NewURLParser().Parse(event)

	assert.Equal(t, meetingApp, "Zoom")
	assert.Equal(t, url, "https://zoom.us/j/12345?pwd=xxx")
}

func TestParseZoomEventUrl(t *testing.T) {
	event := &calendar.Event{
		Description: `──────────<br><br>さんがあなたを予約されたZoomミーティングに招待しています。<br><br>Zoomミーティングに参加する<br><a href="https://zoom.us/j/223344?pwd=yyy">https://zoom.us/j/223344?pwd=yyy</a><br><br>ミーティングID: 123 4567 7890<br>パスコード: 123456<br><br><br>──────────`,
	}
	meetingApp, url := NewURLParser().Parse(event)

	assert.Equal(t, meetingApp, "Zoom")
	assert.Equal(t, url, "https://zoom.us/j/223344?pwd=yyy")
}

func TestParseTeamsEventUrl(t *testing.T) {
	// テストデータ
	event := &calendar.Event{
		Description: `________________________________________________________________________________
Microsoft Teams ミーティング
コンピュータ、モバイルアプリケーション、またはルームデバイスで参加する
ここをクリックして会議に参加してください<https://teams.microsoft.com/l/meetup-join/aaaaaaa>
会議 ID: 123 123 123 123
パスコード: aaaaaa
Teams 
のダウンロード<https://www.microsoft.com/en-us/microsoft-teams/download-app> 
| Web 
に参加<https://www.microsoft.com/microsoft-teams/join-a-meeting>
詳細情報ヘルプ<https://aka.ms/JoinTeamsMeeting> | 
会議のオプション<https://teams.microsoft.com/meetingOptions/?organizerId=ss>
________________________________________________________________________________
`,
	}
	meetingApp, url := NewURLParser().Parse(event)

	assert.Equal(t, meetingApp, "Teams")
	assert.Equal(t, url, "https://teams.microsoft.com/l/meetup-join/aaaaaaa")
}
