package gcal_run

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGCalGetCalendarEvents(t *testing.T) {
	credentialPath := "/tmp/gcal_run_test/credential.json"
	// credentialPathが存在するか確認する
	_, err := os.Stat(credentialPath)
	if err != nil {
		t.Skip("クレデンシャルファイルが存在しないので、テストはスキップ")
	}
	// oauth_tokenが存在するか確認する
	oauthTokenPath := "/tmp/oauth_token"
	_, err = os.Stat(oauthTokenPath)
	if err != nil {
		t.Skip("oauth_tokenファイルが存在しないので、テストはスキップ")
	}
	dateStr := "2023-12-30T11:15:00+09:00"
	// Time.time型に変換する
	basisTime, err := time.Parse(time.RFC3339, dateStr)
	assert.NoError(t, err)
	gcal := NewCalendar(credentialPath, oauthTokenPath)
	calendarEventList, err := gcal.GetCalendarEvents(basisTime)
	assert.NoError(t, err)
	assert.Equal(t, "test1", calendarEventList.Items[0].Summary)
	assert.Equal(t, "test2", calendarEventList.Items[1].Summary)
	assert.Equal(t, "test3", calendarEventList.Items[2].Summary)

}
