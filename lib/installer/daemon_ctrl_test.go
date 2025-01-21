package installer

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDaemonCtrl_GeneratePlistStr(t *testing.T) {
	actual := NewDaemonCtrl().GeneratePlistStr()
	assert.Contains(t, actual, "gcal_run.log")
}

func TestDaemonCtrl_GetPListPath(t *testing.T) {
	assert.Contains(t, NewDaemonCtrl().GetPListPath(), "gcal_run.plist")
}

func TestDaemonCtrl_CreatePListFile(t *testing.T) {
	os.Setenv("GCAL_RUN_TEST", "1")
	daemonCtrl := NewDaemonCtrl()
	err := daemonCtrl.CreatePListFile(false)
	assert.NoError(t, err)
	err = daemonCtrl.DeletePListFile()
	assert.NoError(t, err)
}

func TestDaemonCtrl_StartStopDaemon(t *testing.T) {
	os.Setenv("GCAL_RUN_TEST", "1")
	daemonCtrl := NewDaemonCtrl()

	err := daemonCtrl.CreatePListFile(false)
	assert.NoError(t, err)

	isRunning, err := daemonCtrl.IsDaemonRunning()
	assert.NoError(t, err)
	assert.False(t, isRunning)

	err = daemonCtrl.StartDaemon()
	assert.NoError(t, err)

	isRunning, err = daemonCtrl.IsDaemonRunning()
	assert.NoError(t, err)
	assert.True(t, isRunning)

	err = daemonCtrl.StopDaemon()
	assert.NoError(t, err)

	isRunning, err = daemonCtrl.IsDaemonRunning()
	assert.NoError(t, err)
	assert.False(t, isRunning)

	err = daemonCtrl.DeletePListFile()
	assert.NoError(t, err)

}

func TestDaemonCtrl_IsDaemonRunning(t *testing.T) {
	//os.Setenv("GCAL_RUN_TEST", "1")
	daemonCtrl := NewDaemonCtrl()

	isRunning, err := daemonCtrl.IsDaemonRunning()
	assert.NoError(t, err)
	assert.False(t, isRunning)
}
