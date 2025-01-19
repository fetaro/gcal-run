package main

import (
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/gcal_run"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin/v2"
)

var version string // ビルドスクリプトで埋め込む
var poolingIntervalSec int = 30
var (
	app            = kingpin.New("gcal_run", "gcal_run: GoogleカレンダーTV会議強制起動ツール")
	credentialPath = app.Flag("credential", "GoogleAPIのクレデンシャルファイル").Short('c').Required().ExistingFile()
	installDir     = app.Flag("dir", "インストールディレクトリ").Default(common.DefaultInstallDir()).ExistingDir()
	minuteAgo      = app.Flag("minute", "会議開始の何分前に起動するか").Default(strconv.Itoa(common.DefaultMinutesAgo)).Int()
	browserApp     = app.Flag("browser", "ブラウザアプリケーション").Default(common.DefaultBrowserApp).ExistingDir()
)

func main() {
	app.Version(version)
	app.Parse(os.Args[1:])
	config := common.NewConfig(*credentialPath, *installDir, *minuteAgo, *browserApp)
	runner := gcal_run.NewRunner(config)
	logger := gcal_run.GetLogger()
	logger.Info(config.String())
	// ctrl+cで終了したときのシグナルをキャッチする
	sigs := make(chan os.Signal, 1)
	// 特定のシグナルをキャッチするように設定
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// シグナルを受け取るためのゴルーチンを開始
	go func() {
		sig := <-sigs
		logger.Info("シグナル%sをキャッチ", sig)
		runner.CleanUp()
		logger.Info("終了")
		os.Exit(0)
	}()
	logger.Info("開始。Googleカレンダーのイベントを%s秒毎にチェックします。", poolingIntervalSec)
	for {
		err := runner.Run()
		if err != nil {
			logger.Error("Error: %v", err)
			os.Exit(1)
		}
		logger.Debug("wait %d sec", poolingIntervalSec)
		time.Sleep(time.Duration(poolingIntervalSec) * time.Second)
	}
}
