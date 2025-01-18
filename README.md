GoogleカレンダーWeb会議自動起動ツール(gcal-run)
======================

### これは何？

GoogleカレンダーにあるWeb会議のイベントを定期的にチェックし、開始前に自動的にブラウザで起動するツールです。

Web会議の開始時間を忘れがちな人にオススメです。

## 説明

* インストールするとMacOSのバックグラウンドで動作します。再起動後も自動的に動作開始します。
* Googleカレンダーをチェックし、2分前に開始するGoogleMeet、Zoom、もしくはMicrosoft Teamsの会議のURLがある場合、そのURLをブラウザで開きます。
* 一度起動した会議は内部に記録しておき、再度起動しないようになっています。
* 30秒ごとにカレンダーをチェックします。
* 「何分前にブラウザで起動するか」や「起動するブラウザ」はインストール時の設定で変更することができます。

### 対象となるカレンダーイベントの条件

カレンダーのイベントがオンライン会議かとうか判定は以下の通り

判定条件

* 条件１：オンラインミーティングの設定がされていてボタンが有る場合、設定されたURLを採用
* 条件２：会議の場所に対象のURL(※)が含まれる場合それを採用
* 条件３：本文に対象のURL(※)が含まれる場合それを採用

 ![image](doc/1.png)

(※) 対象のURL

* https://meet.google.com
* https://zoom.us
* https://teams.microsoft.com/l/meetup-join


## インストール方法

### 1. GoogleAPIのクレデンシャルファイルの取得

[このマニュアル](https://github.com/fetaro/gcal-run/wiki/how_to_get_google_calendar_api_credential_file)に従い
Google Calendar APIのJSON形式のクレデンシャルファイルを取得してください。

### 2. ツールのインストール

[ダウンロードページ](https://github.com/fetaro/gcal-run/releases)から最新のバイナリをダウンロードしてください。

* CPUがAppleシリコンの場合(M1など)は `gcal-run_darwin_arm64_x.x.x.tar.gz`
* CPUがIntelの場合は `gcal-run_darwin_amd64_x.x.x.tar.gz`

解凍して、以下のコマンドでインストーラを実行してください。
```text
# 解答したディレクトリに移動
cd 
# インストラーの実行
./installer (クレデンシャルファイルのパス)
```

ここで、「アプリが悪質なソフトウェアであるかどうかAppleで確認できない」と表示された場合は、
アップルメニュー  ＞「システム設定」と選択し、サイドバーで「プライバシーとセキュリティ」を選択
`installerは発行元が不明なため、使用をブロックしました` というメッセージの箇所を許可してください。

プロンプトの指示に従ってインストールしてください。手動で実行する場合はインストールディレクトリの指定が必要であるため、忘れないように記録しておいてください。

### 手動で実行して使う場合

#### 起動

以下のコマンドでプログラムを起動してください。このプログラムが起動している間は自動で会議が始まります
```text
./gcal_run --dir (インストールディレクトリ) --credential (クレデンシャルファイル)
```

ここで、「アプリが悪質なソフトウェアであるかどうかAppleで確認できない」と表示された場合は、
アップルメニュー  ＞「システム設定」と選択し、サイドバーで「プライバシーとセキュリティ」を選択
`gcal_runは発行元が不明なため、使用をブロックしました` というメッセージの箇所を許可してください。

#### 停止

ctrl+cをおしてプログラムを停止してください。

### 常駐プログラムとして使う場合

#### 起動

以下のコマンドでバックグラウンドプロセスを起動してください。
```text
launchctl load ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

ここで、「アプリが悪質なソフトウェアであるかどうかAppleで確認できない」と表示された場合は、
アップルメニュー  ＞「システム設定」と選択し、サイドバーで「プライバシーとセキュリティ」を選択
`gcal_runは発行元が不明なため、使用をブロックしました` というメッセージの箇所を許可してください。

ログは `(インストールディレクトリ)/gcal_run.log` に出力されます。

#### 停止

バックグラウンドプロセスを終了する場合は以下のコマンドを実行してください。
```text
launchctl unload ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

### アンインストール

```text
# 停止
launchctl unload ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist

# 削除
rm ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
rm -rf (インストールディレクトリ)
```

### 設定の変更
バックグラウンドプロセスを停止した後、`${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist`を編集し、再度起動してください。





