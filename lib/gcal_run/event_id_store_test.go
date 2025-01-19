package gcal_run

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventIDStore(t *testing.T) {
	dir, err := os.MkdirTemp(OsUserCacheDir(), "gcal_forcerun_test")
	fileName := time.Now().Format("20060102150405")
	assert.NoError(t, err)
	filePath := dir + "/" + fileName + ".txt"
	db := NewEventIDStore(filePath)

	err = db.SaveID("123")
	assert.NoError(t, err)

	idList, err := db.LoadIDList()
	assert.NoError(t, err)
	assert.Equal(t, []string{"123"}, idList)

	err = db.SaveID("456")
	assert.NoError(t, err)

	err = db.SaveID("")
	assert.Error(t, err)

	idList, err = db.LoadIDList()
	assert.NoError(t, err)
	assert.Equal(t, []string{"123", "456"}, idList)

	isInclude, err := db.IsInclude("123")
	assert.NoError(t, err)
	assert.True(t, isInclude)

	isInclude, err = db.IsInclude("1234")
	assert.NoError(t, err)
	assert.False(t, isInclude)
}

func TestNewEventIDStoreFail(t *testing.T) {
	filePath := "/does_not_exist_dir/hoge.txt"
	db := NewEventIDStore(filePath)
	err := db.SaveID("123")
	assert.Error(t, err)
}
