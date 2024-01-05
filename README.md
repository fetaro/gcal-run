Googleカレンダーイベント自動起動ツール(gcal-run)
======================

### これは何？

GoogleカレンダーにあるWeb会議のイベントを定期的にチェックし、開始前に自動的にブラウザで起動するツールです。

Web会議の開始時間を忘れがちな人にオススメです。

## 詳しい説明

* インストールするとMacOSのバックグラウンドで動作します。具体的にはMacOSのlaunchdを用いてバックグラウンドサービスとして登録するため、再起動しても自動的に起動します。
* 毎時13,14,28,29,43,44,58,59分にGoogleカレンダーをチェックします。
* 2分以内のカレンダーイベントにGoogleMeet、 Zoom、 もしくはMicrosoft Teamsの会議のURLがある場合、そのURLをブラウザで自動的に起動します。
* 一度起動した会議は記録しておき、再度起動しないようになっています。
* 「何分前に起動するか」や「起動するブラウザ」はインストール時の設定で変更することができます。

## インストール方法

### 1. GoogleAPIのクレデンシャルファイルの取得

[このマニュアル](https://developers.google.com/calendar/api/guides/overview?hl=ja)に従いGoogle Calendar APIのクレデンシャルファイルを取得してください。
クレデンシャルファイルはJSON形式です。

### 2. ツールのインストール

[ダウンロードページ](https://github.com/fetaro/gcal-run/releases)から最新のバイナリをダウンロードしてください。

* M1 Macの場合は `gcal-run_darwin_arm64_x.x.x.tar.gz`
* Intel Macの場合は `gcal-run_darwin_amd64_x.x.x.tar.gz`

解凍して、以下のコマンドでインストーラを実行してください。
```text
$ ./bin/installer (クレデンシャルファイルのパス)
```

プロンプトの指示に従ってインストールしてください。

## 使い方

### 起動

以下のコマンドでバックグラウンドプロセスを起動してください。
```text
$ launchctl load ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

ログは `(インストールディレクトリ)/gcal_run.log` に出力されます。

### 停止

バックグラウンドプロセスを終了する場合は以下のコマンドを実行してください。
```text
$ launchctl unload ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

### アンインストール

```text
# 停止
$ launchctl unload ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist

# 削除
$ rm ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
$ rm -rf (インストールディレクトリ)
```

### 設定の変更
バックグラウンドプロセスを停止した後、`${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist`を編集し、再度起動してください。

