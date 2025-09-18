// 強制起動に失敗したイベントの原因を探るために、イベントのパースに失敗したイベントを印字する
package main

import (
	"os"

	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/gcal_run"

	"github.com/alecthomas/kingpin/v2"
)

var version string // ビルドスクリプトで埋め込む
var (
	app        = kingpin.New("gcal_run", "gcal_run: "+common.ToolName)
	configPath = app.Flag("config", "設定ファイルのパス").Short('c').Default(common.GetConfigPath(common.GetAppDir())).String()
)

func main() {
	app.Version(version)
	app.Parse(os.Args[1:])
	config, err := common.LoadConfigFromPath(*configPath)
	if err != nil {
		panic(err)
	}
	runner := gcal_run.NewRunner(config, common.GetAppDir())
	logger, err := gcal_run.GetLogger(common.GetLogPath(common.GetAppDir()))
	if err != nil {
		panic(err)
	}
	err = runner.ParserDebugRun()
	if err != nil {
		logger.Error("エラーのため異常終了: %v", err)
		os.Exit(1)
	}
}
