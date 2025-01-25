GoogleカレンダーWeb会議自動起動ツール(gcal-run)
======================

### これは何？

GoogleカレンダーにあるWeb会議のイベントを定期的にチェックし、開始前に自動的にブラウザで起動するツールです。

Web会議の開始時間を忘れがちな人にオススメです。

## 説明

* インストールするとMacOSのLaunchAgentsの機能を用いて常駐プロセスとして動作します。OS再起動後も自動的に動作開始します。
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

[このマニュアル](https://github.com/fetaro/gcal-run/wiki/how_to_get_google_calendar_api_credential_file)などを参考にして
Google Calendar APIのJSON形式のクレデンシャルファイルを取得してください。

社内の方は私に依頼いただいたら用意できます。

### 2. ツールのインストール

[ダウンロードページ](https://github.com/fetaro/gcal-run/releases)から最新のバイナリをダウンロードしてください。

* CPUがAppleシリコンの場合(M1など)は `gcal-run_darwin_arm64_x.x.x.tar.gz`
* CPUがIntelの場合は `gcal-run_darwin_amd64_x.x.x.tar.gz`

ダウンロードしたファイルを解凍してください。Finderでダブルクリックすれば解答できます。

ターミナルを起動してください。

解答したディレクトリに移動してください
```bash
cd (ダウンロードしたディレクトリ)
```

例：ダウンロードしたディレクトリが「ダウンロード/gcal-run_darwin_amd64_v2.0.5」の場合

```bash
cd ダウンロード/gcal-run_darwin_amd64_v2.0.5
```

以下のコマンドでインストラーを実行してください

```bash
./installer install
```

指示に従ってインストールしてください

これで `${HOME}/.gcal_run/` に必要なファイルがインストールされます。


ここで、「アプリが悪質なソフトウェアであるかどうかAppleで確認できない」と表示された場合は、
アップルメニュー  ＞「システム設定」と選択し、サイドバーで「プライバシーとセキュリティ」を選択
`installerは発行元が不明なため、使用をブロックしました` というメッセージの箇所を許可してください。

## 実行方法

### 常駐プログラムとして使う場合

#### 起動

以下のコマンドで常駐プロセスを起動してください。
```text
launchctl load ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

ここで、「アプリが悪質なソフトウェアであるかどうかAppleで確認できない」と表示された場合は、
アップルメニュー  ＞「システム設定」と選択し、サイドバーで「プライバシーとセキュリティ」を選択
`gcal_runは発行元が不明なため、使用をブロックしました` というメッセージの箇所を許可してください。

#### 常駐プロセスが動いていることの確認
プロセスの確認
```
ps -ef | grep gcal | grep -v grep
# 行が表示されたらプロセスは起動している
```

もしくは、LaunchAgentsのコマンドで確認
```
launchctl list | grep gcal
# 一番左の文字が「-」ではなく、数字が表示されている場合は、プロセスが起動している
```

#### ログの確認
```
tail -f ${HOME}/.gcal_run/gcal_run.log
```

#### 停止

常駐プロセスを終了する場合は以下のコマンドを実行してください。
```text
launchctl unload ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

#### 設定の変更
常駐プロセスを停止した後、
`${HOME}/.gcal_run/config.json`を編集し、
再度常駐プロセスを起動してください。


#### アンインストール

インストールディレクトリで以下を実行
```text
./installer uninstall 
```


#### アップデート
インストールディレクトリで以下を実行
```text
./installer update 
```


### 手動で実行して使う場合

#### 起動

以下のコマンドでプログラムを起動してください。このプログラムが起動している間は自動で会議が始まります
```text
cd ${HOME}/.gcal_run
./gcal_run
```

#### 停止

ctrl+cをおしてプログラムを停止してください。

## v1.x.xからv2.x.xへのアップデート手順

v1.x.x系はuninstallコマンドがないので、手動でのツールの削除が必要です。

以下の手順でv1.x.xのツールを削除してください。
```text
# 常駐プロセスの停止
launchctl unload ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist

# 常駐プロセス定義ファイルの削除
rm ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist

# ツールの削除
# デフォルトのインストール先は ${HOME}/.gcal_run です
rm -rf (インストールディレクトリ)

```

[ダウンロードページ](https://github.com/fetaro/gcal-run/releases)から最新のバイナリをダウンロードしてください。

* CPUがAppleシリコンの場合(M1など)は `gcal-run_darwin_arm64_x.x.x.tar.gz`
* CPUがIntelの場合は `gcal-run_darwin_amd64_x.x.x.tar.gz`

解凍して、以下のコマンドでインストーラを実行してください。
```text
# 解凍したディレクトリに移動
cd 
# インストラーの実行
./installer install
```

## その他

- [リリースノート](RELEASE_NOTE.md)
