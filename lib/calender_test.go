package lib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// このテストは/tmp/gcal_run_test/credential.jsonと/tmp/gcal_run_test/oauth_tokenが存在する場合のみ実行される
func TestGCalGetCalenderEvents(t *testing.T) {
	config := NewConfig("/tmp/gcal_run_test/credential.json", "/tmp/gcal_run_test", 30, "/Applications/Google Chrome.app")
	err := config.IsValid()
	assert.NoError(t, err)
	dateStr := "2023-12-30T11:15:00+09:00"
	// Time.time型に変換する
	basisTime, err := time.Parse(time.RFC3339, dateStr)
	assert.NoError(t, err)
	gcal := NewCalender(config)
	calenderEventList, err := gcal.GetCalenderEvents(basisTime)
	assert.NoError(t, err)
	assert.Equal(t, "test1", calenderEventList.Items[0].Summary)
	assert.Equal(t, "test2", calenderEventList.Items[1].Summary)
	assert.Equal(t, "test3", calenderEventList.Items[2].Summary)
}
