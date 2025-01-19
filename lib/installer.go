package lib

import (
	"fmt"
	"os"
	"time"
)

type Installer struct {
}

func NewInstaller() *Installer {
	return &Installer{}
}

func (i *Installer) Install(c *Config) error {
	// c.InstallDirを作成する
	err := os.MkdirAll(c.InstallDir, 0755)
	if err != nil {
		return fmt.Errorf("インストールディレクトリの作成に失敗: %v", err)
	}
	tokenGetter := NewOAuthTokenGetter()
	_, err = tokenGetter.GetAndSaveToken(c.CredentialPath, c.TokenPath, c.BrowserApp)
	calendar := NewCalendar(c)
	events, err := calendar.GetCalendarEvents(time.Now())
	if err != nil {
		return fmt.Errorf("カレンダーサービスのアクセスに失敗しました: %v", err)
	} else {
		fmt.Println("カレンダーサービスのアクセスに成功しました")
		fmt.Println("取得できたカレンダーイベント：")
		for _, item := range events.Items {
			fmt.Println(item.Summary)
		}
	}

	// バイナリをコピー
	src := "./gcal_run"
	fmt.Printf("バイナリをコピーします: %s -> %s\n", src, c.BinPath)
	err = CopyFile(src, c.BinPath)
	if err != nil {
		return fmt.Errorf("バイナリのコピーに失敗: %v", err)
	}

	fmt.Println("バイナリに実行権限を付与します")
	err = os.Chmod(c.BinPath, 0755)
	if err != nil {
		return fmt.Errorf("バイナリに実行権限を付与できませんでした: %v", err)
	}
	fmt.Printf("バイナリに実行権限を付与しました: %s\n", c.BinPath)

	// plistファイルを作成
	daemonCtrl := NewDaemonCtrl()
	err = daemonCtrl.CreatePListFile(c)
	if err != nil {
		return err
	}

	fmt.Printf(`
============================================
インストールが完了しました。
`)
	fmt.Printf(`
## 設定
クレデンシャルファイルパス： %s
利用ブラウザ : %s
会議起動時間 : %d 分前
`, c.CredentialPath, c.BrowserApp, c.MinutesAgo,
	)

	fmt.Printf(`
## インストールしたファイル
インストールディレクトリ : %s
常駐プロセス(LaunchAgents)ファイル : %s
`, c.InstallDir, daemonCtrl.GetPListPath())

	fmt.Printf(`
## 使い方
### 常駐プロセスの起動
$ launchctl load %s
### 常駐プロセスの停止
$ launchctl unload %s
### ログの確認
$ tail -f %s
`, daemonCtrl.GetPListPath(), daemonCtrl.GetPListPath(), c.LogPath)

	return nil
}
