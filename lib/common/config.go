package common

import (
	"fmt"
	"os"
	"path"
)

type Config struct {
	CredentialPath   string
	InstallDir       string
	MinutesAgo       int
	BrowserApp       string
	BinPath          string
	LogPath          string
	TokenPath        string
	EventIDStorePath string
}

func NewConfig(credentialPath string, installDir string, minutesAgo int, browserApp string) *Config {
	return &Config{
		CredentialPath:   credentialPath,
		InstallDir:       installDir,
		MinutesAgo:       minutesAgo,
		BrowserApp:       browserApp,
		BinPath:          path.Join(installDir, "gcal_run"),
		LogPath:          path.Join(installDir, "gcal_run.log"),
		TokenPath:        path.Join(installDir, "oauth_token"),
		EventIDStorePath: path.Join(installDir, "event_id_store"),
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

func (c *Config) String() string {
	return fmt.Sprintf(`設定値:クレデンシャルファイルパス=%s,会議起動時間=%d 分前,ブラウザアプリケーション=%s`, c.CredentialPath, c.MinutesAgo, c.BrowserApp)
}
