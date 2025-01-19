package installer

import (
	"bufio"
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/gcal_run"
	"os"
	"strconv"
	"time"
)

type Installer struct {
}

func NewInstaller() *Installer {
	return &Installer{}
}
func (i *Installer) ScanInput(credentialPath string) (*common.Config, error) {
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	fmt.Printf("インストール先ディレクトリを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", common.DefaultInstallDir())
	scanner.Scan()
	installDir := scanner.Text()
	if installDir == "" {
		installDir = common.DefaultInstallDir()
	}
	// installDirが存在しない場合は作る
	if _, err := os.Stat(installDir); os.IsNotExist(err) {
		err := os.MkdirAll(installDir, 0755)
		if err != nil {
			fmt.Printf("ディレクトリを作成できませんでした: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("ディレクトリを作成しました: %s\n", installDir)
	} else {
		fmt.Printf("ディレクトリが既に存在します。: %s\n", installDir)
		fmt.Printf("中身を空にして、インストールしますか？ (y/n) > ")
		scanner2 := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
		scanner2.Scan()
		yOrN := scanner2.Text()
		if yOrN == "y" {
			// installDirの中身を空にする
			err := os.RemoveAll(installDir)
			if err != nil {
				fmt.Printf("ディレクトリを空にできませんでした: %v\n", err)
				os.Exit(1)
			} else {
				err := os.MkdirAll(installDir, 0755)
				if err != nil {
					fmt.Printf("ディレクトリを作成できませんでした: %v\n", err)
					os.Exit(1)
				}
				fmt.Printf("ディレクトリを空にして再作成しました: %s\n", installDir)
			}
		} else {
			fmt.Println("インストールを中止します")
			os.Exit(1)
		}
	}

	var browserApp string
	for {
		fmt.Printf("ブラウザアプリケーションのパスを指定してください\nデフォルトは「%s」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", common.DefaultBrowserApp)
		scanner.Scan()
		browserApp = scanner.Text()
		_, err := os.Stat(browserApp)
		if browserApp == "" {
			browserApp = common.DefaultBrowserApp
			break
		}
		if os.IsNotExist(err) {
			fmt.Println("ブラウザアプリケーションが存在しません。再度入力してください")
		} else {
			break
		}
	}

	var minutesAgoStr string
	var minutesAgo int
	for {
		fmt.Printf("会議の何分前に起動するか指定してください\nデフォルトは「%d分」です。デフォルトで良い場合は何も入力せずにEnterを押してください\n> ", common.DefaultMinutesAgo)
		scanner.Scan()
		minutesAgoStr = scanner.Text()
		_, err := os.Stat(minutesAgoStr)
		if minutesAgoStr == "" {
			minutesAgo = common.DefaultMinutesAgo
			break
		}
		minutesAgo, err = strconv.Atoi(minutesAgoStr)
		if err != nil {
			fmt.Println("数値を入力してください")
			continue
		} else {
			break
		}
	}
	config := common.NewConfig(credentialPath, installDir, minutesAgo, browserApp)
	return config, nil
}

func (i *Installer) Install(c *common.Config) error {
	fmt.Println("ツールをダウンロードし、インストールディレクトリに展開します")
	latestVersion, err := NewGithubService().GetLatestVersion()
	if err != nil {
		return err
	}
	NewDownloader().Download(latestVersion, c.InstallDir)
	fmt.Printf("ツールをダウンロードしました: version=%s, path=%s\n", latestVersion, c.BinPath)

	fmt.Println("バイナリに実行権限を付与します")
	err = os.Chmod(c.BinPath, 0755)
	if err != nil {
		return fmt.Errorf("バイナリに実行権限を付与できませんでした: %v", err)
	}
	fmt.Printf("バイナリに実行権限を付与しました: %s\n", c.BinPath)

	// トークンの取得
	tokenGetter := gcal_run.NewOAuthTokenGetter()
	_, err = tokenGetter.GetAndSaveToken(c.CredentialPath, c.TokenPath, c.BrowserApp)
	// カレンダーとの疎通確認
	calendar := gcal_run.NewCalendar(c)
	events, err := calendar.GetCalendarEvents(time.Now())
	if err != nil {
		return fmt.Errorf("カレンダーサービスのアクセスに失敗しました: %v", err)
	} else {
		fmt.Println("Googleカレンダーとの疎通に成功しました")
		fmt.Println("取得できたカレンダーイベント：")
		for index, item := range events.Items {
			fmt.Printf("- %d : %s\n", index, item.Summary)
		}
	}

	// plistファイルを作成
	daemonCtrl := NewDaemonCtrl()
	err = daemonCtrl.CreatePListFile(c)
	if err != nil {
		return err
	}

	fmt.Printf(`
============================================
インストールが完了しました。

インストールディレクトリ : %s
常駐プロセス(LaunchAgents)ファイル : %s
`, c.InstallDir, daemonCtrl.GetPListPath())

	return nil
}
