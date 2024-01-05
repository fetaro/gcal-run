package lib

import (
	"fmt"
	"os"
	"path"
)

const (
	DefaultBrowserApp    = "/Applications/Google Chrome.app"
	DefaultSearchMinutes = 2
)

type Config struct {
	CredentialPath   string
	InstallDir       string
	SearchMinutes    int
	BrowserApp       string
	BinPath          string
	LogPath          string
	TokenPath        string
	EventIDStorePath string
	PlistPath        string
}

func NewConfig(credentialPath string, installDir string, searchMinutes int, browserApp string) *Config {
	return &Config{
		CredentialPath:   credentialPath,
		InstallDir:       installDir,
		SearchMinutes:    searchMinutes,
		BrowserApp:       browserApp,
		BinPath:          path.Join(installDir, "gcal_run"),
		LogPath:          path.Join(installDir, "gcal_run.log"),
		TokenPath:        path.Join(installDir, "oauth_token"),
		EventIDStorePath: path.Join(installDir, "event_id_store"),
		PlistPath:        path.Join(os.Getenv("HOME"), "Library/LaunchAgents/com.github.fetaro.gcal_run.plist"),
	}
}

func (c *Config) IsValid() error {
	if _, err := os.Stat(c.CredentialPath); os.IsNotExist(err) {
		return fmt.Errorf("クレデンシャルファイルを読み取れません: %v", err)
	}
	if _, err := os.Stat(c.InstallDir); os.IsNotExist(err) {
		return fmt.Errorf("インストール先ディレクトリが存在しません: %v", err)
	}
	return nil
}

func (c *Config) InstructString() string {
	return fmt.Sprintf(`
## 設定
クレデンシャルファイルパス： %s
利用ブラウザ : %s
会議起動時間 : %d 分前

## インストールしたファイル
インストール先 : %s
LaunchAgentファイル : %s

## 使い方
以下のコマンドでサービスデーモンを起動してください。
$ launchctl load %s

サービスデーモンを終了する場合は以下のコマンドを実行してください。
$ launchctl unload %s

ログは以下の場所に出力されます。
%s

アンインストールする場合はデーモンを終了した後、以下のコマンドでインストールディレクトリを削除してください。
$ rm -rf %s
`,
		c.CredentialPath, c.BrowserApp, c.SearchMinutes,
		c.InstallDir, c.PlistPath,
		c.PlistPath, c.PlistPath, c.LogPath, c.InstallDir)
}

func (c *Config) GeneratePlistStr() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>RunAtLoad</key>
	<false/>

	<key>KeepAlive</key>
	<true/>

	<key>Label</key>
	<string>com.github.fetaro.gcal_run</string>

	<key>ProgramArguments</key>
	<array>
		<string>%s</string>
        <string>--credential</string>
		<string>%s</string>
        <string>--dir</string>
		<string>%s</string>
        <string>--minute</string>
		<string>%d</string>
        <string>--browser</string>
		<string>%s</string>
	</array>

	<key>StandardErrorPath</key>
	<string>%s</string>

	<key>StandardOutPath</key>
	<string>%s</string>
</dict>
</plist>
`, c.BinPath, c.CredentialPath, c.InstallDir, c.SearchMinutes, c.BrowserApp, c.LogPath, c.LogPath)
}
