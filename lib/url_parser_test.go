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
	url, err := NewURLParser().Parse(event)

	assert.NoError(t, err)
	assert.Equal(t, "https://meet.google.com/xxx", url)
}

func TestParseZoomEventUrlWhenLocationHasUrl(t *testing.T) {
	event := &calendar.Event{
		Location: "https://zoom.us/j/xxx",
	}
	url, err := NewURLParser().Parse(event)

	assert.NoError(t, err)
	assert.Equal(t, "https://zoom.us/j/xxx", url)
}

func TestParseZoomEventUrlWhenDescriptionHasUrl(t *testing.T) {
	event := &calendar.Event{
		Description: "<a href=\"https://zoom.us/j/12345?pwd=xxx\"><u>https://zoom.us/j/12345?pwd=xxx</u></a>",
	}
	url, err := NewURLParser().Parse(event)

	assert.NoError(t, err)
	assert.Equal(t, "https://zoom.us/j/12345?pwd=xxx", url)
}

func TestParseZoomEventUrl(t *testing.T) {
	event := &calendar.Event{
		Description: `──────────<br><br>さんがあなたを予約されたZoomミーティングに招待しています。<br><br>Zoomミーティングに参加する<br><a href="https://zoom.us/j/223344?pwd=yyy">https://zoom.us/j/223344?pwd=yyy</a><br><br>ミーティングID: 123 4567 7890<br>パスコード: 123456<br><br><br>──────────`,
	}
	url, err := NewURLParser().Parse(event)

	assert.NoError(t, err)
	assert.Equal(t, "https://zoom.us/j/223344?pwd=yyy", url)
}

func TestParseZoomEventUrl2(t *testing.T) {
	event := &calendar.Event{
		Description: `Zoomミーティングに参加する
https://zoom.us/j/1234?pwd=12345`,
	}
	url, err := NewURLParser().Parse(event)

	assert.NoError(t, err)
	assert.Equal(t, "https://zoom.us/j/1234?pwd=12345", url)
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
	url, err := NewURLParser().Parse(event)

	assert.NoError(t, err)
	assert.Equal(t, "https://teams.microsoft.com/l/meetup-join/aaaaaaa", url)
}

func TestParseTeamsEventUrlWhenDescriptionComplexUrl(t *testing.T) {
	// テストデータ
	event := &calendar.Event{
		Description: `
Microsoft Teams ミーティング
コンピュータ、モバイルアプリケーション、またはルームデバイスで参加する
ここをクリックして会議に参加してください<https://teams.microsoft.com/l/meetup-join/19%3ameeting_OGhread.v2/0?context=%a%b-c-d>
会議 ID: 123 2345 345 456
パスコード: xxxxx
Teams のダウンロード<https://www.microsoft.com/en-us/microsoft-teams/download-app> | Web に参加<https://www.microsoft.com/microsoft-teams/join-a-meeting>
詳細情報ヘルプ<https://aka.ms/JoinTeamsMeeting> | 会議のオプション<https://teams.microsoft.com/meetingOptions/?organizerId=123-234&tenantId=43&threadId=1_m@thread.v2&messageId=0&language=ja-JP>
`,
	}
	url, err := NewURLParser().Parse(event)
	assert.NoError(t, err)
	assert.Equal(t, "https://teams.microsoft.com/l/meetup-join/19%3ameeting_OGhread.v2/0?context=%a%b-c-d", url)
}
