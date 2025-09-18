package gcal_run

import (
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGCalGetCalendarEvents(t *testing.T) {
	credentialPath := "/tmp/gcal_run_test/credential.json"
	if !common.FileExists(credentialPath) {
		t.Skip("クレデンシャルファイルが存在しないので、テストはスキップ")
	}
	oauthTokenPath := "/tmp/oauth_token"
	if !common.FileExists(oauthTokenPath) {
		t.Skip("oauth_tokenファイルが存在しないので、テストはスキップ")
	}
	dateStr := "2023-12-30T11:15:00+09:00"
	// Time.time型に変換する
	basisTime, err := time.Parse(time.RFC3339, dateStr)
	assert.NoError(t, err)
	gcal := NewCalendar(credentialPath, oauthTokenPath)
	calendarEventList, err := gcal.GetCalendarEvents(basisTime,4)
	assert.NoError(t, err)
	assert.Equal(t, "test1", calendarEventList.Items[0].Summary)
	assert.Equal(t, "test2", calendarEventList.Items[1].Summary)
	assert.Equal(t, "test3", calendarEventList.Items[2].Summary)

}
