// このコードは、以下のサイトのコードを参考にしています。
// https://github.com/googleapis/google-api-go-client/blob/main/examples/main.go

package installer

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"github.com/fetaro/gcal_forcerun_go/lib/gcal_run"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type OAuthTokenGetter struct {
	ConfirmOverwrite bool
}

func NewOAuthTokenGetter(confirmOverwrite bool) *OAuthTokenGetter {
	return &OAuthTokenGetter{ConfirmOverwrite: confirmOverwrite}
}

func (o *OAuthTokenGetter) saveToken(file string, token *oauth2.Token) error {
	logger, err := gcal_run.GetLogger("")
	if err != nil {
		return err
	}
	if common.FileExists(file) {
		if PrintAndScanStdInput("トークンファイルが既に存在します。上書きしますか？ (y/n) > ") != "y" {
			logger.Info("トークンファイルの上書きをキャンセルしました。")
			return nil
		}
	}
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("トークンファイルの作成に失敗しました。 %v", err)
	}
	defer f.Close()
	gob.NewEncoder(f).Encode(token)
	logger.Info("OAuthトークンをファイルに保存。トークンファイル=%s", file)
	return nil
}

func (o *OAuthTokenGetter) getTokenFromWeb(credentialPath string, browserApp string) (*oauth2.Token, error) {
	logger, err := gcal_run.GetLogger("")
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	b, err := os.ReadFile(credentialPath)
	if err != nil {
		return nil, fmt.Errorf("fail to read credential file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("fail to parse credential file to config: %v", err)
	}
	ch := make(chan string)
	randState := fmt.Sprintf("st%d", time.Now().UnixNano())
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.Error(rw, "", 404)
			return
		}
		if req.FormValue("state") != randState {
			logger.Error("State doesn't match: req = %#v", req)
			http.Error(rw, "", 500)
			return
		}
		if code := req.FormValue("code"); code != "" {
			fmt.Fprintf(rw, "アプリケーションの認証成功")
			rw.(http.Flusher).Flush()
			ch <- code
			return
		}
		logger.Error("no code")
		http.Error(rw, "", 500)
	}))
	defer ts.Close()

	config.RedirectURL = ts.URL
	authURL := config.AuthCodeURL(randState)
	common.OpenUrl(browserApp, authURL)
	fmt.Printf("ブラウザを使ってこのアプリケーションを認証してください。URL = %s\n", authURL)
	code := <-ch
	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Token exchange error: %v", err)
	}
	return token, nil
}

func (o *OAuthTokenGetter) GetAndSaveToken(credentialPath string, tokenFilePath string, browserApp string) (*oauth2.Token, error) {
	token, err := o.getTokenFromWeb(credentialPath, browserApp)
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %v", err)
	}
	err = o.saveToken(tokenFilePath, token)
	if err != nil {
		return nil, fmt.Errorf("failed to save token: %v", err)
	}
	return token, nil
}
