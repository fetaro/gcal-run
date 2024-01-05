// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// Simple service that only works by printing a log message every few seconds.
package main

import (
	"os"
	"strconv"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/fetaro/gcal_forcerun_go/lib"
)

var (
	credentialPath = kingpin.Flag("credential", "GoogleAPIのクレデンシャルファイル").Required().String()
	installDir     = kingpin.Flag("dir", "インストールディレクトリ").Required().String()
	searchMinute   = kingpin.Flag("minute", "会議開始の何分前に起動するか").Default(strconv.Itoa(lib.DefaultSearchMinutes)).Int()
	browserApp     = kingpin.Flag("browser", "ブラウザアプリケーション").Default(lib.DefaultBrowserApp).String()
)

func main() {
	kingpin.Parse()
	config := lib.NewConfig(*credentialPath, *installDir, *searchMinute, *browserApp)
	runner := lib.NewRunner(config)
	logger := lib.GetLogger()
	logger.Info("開始")
	logger.Info(config.String())
	logger.Info("毎時 13,14,28,29,43,44,58,59 分にカレンダーチェック")
	for {
		// 現在時刻を取得
		now := time.Now()
		// 分が13,14,28,29,43,44,58,59のときに実行
		if now.Minute()%15 == 14 || now.Minute()%15 == 13 {
			err := runner.Run()
			if err != nil {
				logger.Error("Error: %v", err)
				os.Exit(1)
			}
		}
		logger.Debug("wait 1 minutes")
		time.Sleep(60 * time.Second)
	}
}
