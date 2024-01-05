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
	calender := NewCalender(c)
	_, err = calender.GetCalenderEvents(time.Now())
	if err != nil {
		return fmt.Errorf("カレンダーサービスのアクセスに失敗しました: %v", err)
	}

	// plistファイルを作成
	err = os.WriteFile(c.PlistPath, []byte(c.GeneratePlistStr()), 0644)
	if err != nil {
		return fmt.Errorf("サービスデーモンファイル(plist)の作成に失敗: %v", err)
	}
	fmt.Printf("Macのサービスデーモンファイル(plist)を作成しました: %s\n", c.PlistPath)

	// バイナリをコピー
	src := "./gcal_run"
	fmt.Printf("バイナリをコピーします: %s -> %s\n", src, c.BinPath)
	err = CopyFile(src, c.BinPath)
	if err != nil {
		return fmt.Errorf("バイナリのコピーに失敗: %v", err)
	}

	err = os.Chmod(c.BinPath, 0755)
	if err != nil {
		return fmt.Errorf("バイナリに実行権限を付与できませんでした: %v", err)
	}
	fmt.Printf("バイナリに実行権限を付与しました: %s\n", c.BinPath)

	fmt.Println(`============================================
インストールが完了しました。
` + c.InstructString())
	return nil
}
