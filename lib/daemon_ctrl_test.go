package lib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDaemonCtrl_GeneratePlistStr(t *testing.T) {
	config := NewConfig("credpath", "installDir", 30, "/Applications/Google Chrome.app")
	actual := NewDaemonCtrl().GeneratePlistStr(config)
	assert.Contains(t, actual, "credpath")
	assert.Contains(t, actual, "installDir/gcal_run")
	assert.Contains(t, actual, "installDir/gcal_run.log")
	assert.Contains(t, actual, "/Applications/Google Chrome.app")
	assert.Contains(t, actual, "30")
}

func TestDaemonCtrl_GetPListPath(t *testing.T) {
	assert.Contains(t, NewDaemonCtrl().GetPListPath(), "gcal_run.plist")
}

func TestDaemonCtrl_CreatePListFile(t *testing.T) {
	os.Setenv("GCAL_RUN_TEST", "1")
	c := NewConfig("/tmp/gcal_forcerun_secret.json", "/tmp/gcal_run_test", 2, "/Applications/Google Chrome.app")
	daemonCtrl := NewDaemonCtrl()
	err := daemonCtrl.CreatePListFile(c)
	assert.NoError(t, err)
	err = daemonCtrl.DeletePListFile()
	assert.NoError(t, err)
}

func TestDaemonCtrl_StartStopDaemon(t *testing.T) {
	os.Setenv("GCAL_RUN_TEST", "1")
	c := NewConfig("/tmp/gcal_forcerun_secret.json", "/tmp/gcal_run_test", 2, "/Applications/Google Chrome.app")
	daemonCtrl := NewDaemonCtrl()

	err := daemonCtrl.CreatePListFile(c)
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
