package main

import (
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/gcal_run"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin/v2"
)

var version string // ビルドスクリプトで埋め込む
var poolingIntervalSec = 30
var (
	app        = kingpin.New("gcal_run", "gcal_run: "+common.ToolName)
	configPath = app.Flag("config", "設定ファイルのパス").Short('c').Default(common.GetConfigPath(common.GetAppDir())).String()
	debug      = app.Flag("debug", "デバッグログを出力する").Bool()
)

func main() {
	app.Version(version)
	app.Parse(os.Args[1:])
	config, err := common.LoadConfigFromPath(*configPath)
	if err != nil {
		panic(err)
	}
	if *debug {
		// 環境変数にDEBUG=1を設定する
		os.Setenv("DEBUG", "1")
	}
	runner := gcal_run.NewRunner(config, common.GetAppDir())
	logger, err := gcal_run.GetLogger(common.GetLogPath(common.GetAppDir()))
	if err != nil {
		panic(err)
	}
	logger.Info("------------------------------------------------")
	logger.Info("GoogleカレンダーWeb会議自動起動ツール(gcal-run)")
	logger.Info("------------------------------------------------")
	logger.Info("バージョン: %s", version)
	logger.Info("設定ファイルパス: %s", *configPath)
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
		os.Exit(0)
	}()
	// 終了メッセージ
	defer logger.Info("終了")
	logger.Info("開始。Googleカレンダーのイベントを%d秒毎にチェックします。", poolingIntervalSec)
	for {
		err := runner.Run()
		if err != nil {
			logger.Error("エラーのため異常終了: %v", err)
			os.Exit(1)
		}
		logger.Debug("wait %d sec", poolingIntervalSec)
		time.Sleep(time.Duration(poolingIntervalSec) * time.Second)
	}
}
