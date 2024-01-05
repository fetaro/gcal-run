// このコードは、以下のサイトのコードを参考にしています。
// https://github.com/googleapis/google-api-go-client/blob/main/examples/main.go

package lib

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type OAuthTokenGetter struct {
}

func NewOAuthTokenGetter() *OAuthTokenGetter {
	return &OAuthTokenGetter{}
}

func (o *OAuthTokenGetter) saveToken(file string, token *oauth2.Token) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("failed to create token file %v", err)
	}
	defer f.Close()
	gob.NewEncoder(f).Encode(token)
	log.Printf("OAuthトークンをファイルに保存。トークンファイル=%s", file)
	return nil
}

func (o *OAuthTokenGetter) getTokenFromWeb(credentialPath string, browserApp string) (*oauth2.Token, error) {
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
			log.Printf("State doesn't match: req = %#v", req)
			http.Error(rw, "", 500)
			return
		}
		if code := req.FormValue("code"); code != "" {
			fmt.Fprintf(rw, "アプリケーションの認証成功")
			rw.(http.Flusher).Flush()
			ch <- code
			return
		}
		log.Printf("no code")
		http.Error(rw, "", 500)
	}))
	defer ts.Close()

	config.RedirectURL = ts.URL
	authURL := config.AuthCodeURL(randState)
	go OpenUrl(browserApp, authURL)
	log.Printf("ブラウザを使ってこのアプリケーションを認証してください。URL=%s", authURL)
	code := <-ch
	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Token exchange error: %v", err)
	}
	return token, nil
}

func OpenUrl(browserApp, url string) error {
	err := exec.Command("open", "-a", browserApp, url).Run()
	if err != nil {
		return fmt.Errorf("failed to open event url: %v", err)
	}
	return nil
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
