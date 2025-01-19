package installer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubService struct {
}

func NewGithubService() *GithubService {
	return &GithubService{}
}

func (g *GithubService) GetLatestVersion() (*Version, error) {
	url := "https://api.github.com/repos/fetaro/gcal-run/releases"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GitHubのAPIからリリース情報の取得に失敗しました。URL:%s error:%v\n", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GitHubのAPIからリリース情報のリクエストボディの取得に失敗しました。%v\n", err)
	}

	var releases []Release
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return nil, fmt.Errorf("GitHubのAPIからリリース情報のJSONのパースに失敗しました。json=%s, error=%v\n", body, err)
	}
	latestTagName := releases[0].TagName
	return ParseVersionStr(latestTagName)
}
