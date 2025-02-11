# Macへのインストール方法

## インストール方法

### Google Calendar APIのクレデンシャルファイルの取得

Google Calendarから自分の予定を取得するためには、Google Calendar APIの「クレデンシャルファイル」が必要です。

社内の方は管理者に依頼してください。用意します。

その他の方は、[このマニュアル](https://github.com/fetaro/gcal-run/wiki/how_to_get_google_calendar_api_credential_file)などを参考にして Google Calendar APIのJSON形式のクレデンシャルファイルを取得してください。


### ツールのダウンロード

[ダウンロードページ](https://github.com/fetaro/gcal-run/releases)から最新のバイナリをダウンロードしてください。

* CPUがAppleシリコンの場合(M1など)は `gcal-run_darwin_arm64_x.x.x.tar.gz`
* CPUがIntelの場合は `gcal-run_darwin_amd64_x.x.x.tar.gz`

ダウンロードしたファイルを解凍してください。Finderでダブルクリックすれば解答できます。

### ツールのインスト―ル

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
./installer
```

指示に従ってインストールしてください。

注意）ここで、「アプリが悪質なソフトウェアであるかどうかAppleで確認できない」と表示された場合は、
アップルメニュー  ＞「システム設定」と選択し、サイドバーで「プライバシーとセキュリティ」を選択
`installerは発行元が不明なため、使用をブロックしました` というメッセージの箇所を許可してください。

問題なく完了できれば、ホームディレクトリの下の「.gcal_run」というフォルダにに必要なファイルがインストールされて、
常駐プログラムが起動します。


## アンインストール方法

ターミナルで以下を実行し、アンインストールを選択してください
```bash
cd ${HOME}/.gcal_run
./installer 
```


## バージョンアップ方法

ターミナルで以下を実行、バージョンアップを選択してください
```bash
cd ${HOME}/.gcal_run
./installer 
```

## 詳細なマニュアル(エンジニア向け)

### 手動でプログラムを実行する

#### 常駐プロセスの起動

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

#### 常駐プロセスの停止

常駐プロセスを終了する場合は以下のコマンドを実行してください。
```text
launchctl unload ${HOME}/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

#### 設定の変更
常駐プロセスを停止した後、
`${HOME}/.gcal_run/config.json`を編集し、
再度常駐プロセスを起動してください。




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
./installer 
```

## インストールに失敗する場合

[よくある質問](https://github.com/fetaro/gcal-run/wiki)を御覧ください

## その他

- [リリースノート](RELEASE_NOTE.md)
