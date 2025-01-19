package common

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	CredentialPath string
	MinutesAgo     int
	BrowserApp     string
}

func NewConfig(credentialPath string, minutesAgo int, browserApp string) *Config {
	return &Config{
		CredentialPath: credentialPath,
		MinutesAgo:     minutesAgo,
		BrowserApp:     browserApp,
	}
}

func (c *Config) IsValid() error {
	if _, err := os.Stat(c.CredentialPath); os.IsNotExist(err) {
		return fmt.Errorf("クレデンシャルファイルを読み取れません: %v", err)
	}
	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf(`設定値:クレデンシャルファイルパス=%s,会議起動時間=%d 分前,ブラウザアプリケーション=%s`, c.CredentialPath, c.MinutesAgo, c.BrowserApp)
}

func (c *Config) Save() error {
	// JSON形式に変換
	configJSON, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("設定値のJSON変換に失敗しました: %v", err)
	}
	// ファイルに書き込み
	return os.WriteFile(GetConfigPath(GetAppDir()), configJSON, 0644)
}
func LoadConfig() (*Config, error) {
	return LoadConfigFromPath(GetConfigPath(GetAppDir()))
}
func LoadConfigFromPath(configPath string) (*Config, error) {
	configJSON, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("設定ファイルの読み込みに失敗しました。 path=\"%s\"", configPath)
	}
	config := &Config{}
	err = json.Unmarshal(configJSON, config)
	if err != nil {
		return nil, fmt.Errorf("設定ファイルのJSON変換に失敗しました: %v", err)
	}
	return config, nil
}
