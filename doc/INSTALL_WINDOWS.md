# Windowsへのインストール方法

## インストール方法

### Google Calendar APIのクレデンシャルファイルの取得

Google Calendarから自分の予定を取得するためには、Google Calendar APIの「クレデンシャルファイル」が必要です。

社内の方は管理者に依頼してください。用意します。

その他の方は、[このマニュアル](https://github.com/fetaro/gcal-run/wiki/how_to_get_google_calendar_api_credential_file)などを参考にして Google Calendar APIのJSON形式のクレデンシャルファイルを取得してください。


#### ツールのダウンロード

[ダウンロードページ](https://github.com/fetaro/gcal-run/releases)から  `gcal-run_windows_amd64_vx.x.x.zip`  という名前のファイルをダウンロードしてください。x.x.xの部分はバージョンです。2.1.4以降のバージョンをダウンロードしてください。

ダウンロードしたファイルを解凍してください。

解凍すると、以下のようなファイルが展開されます。

* gcal_run.exe : ツール本体
* installer.exe : インストラ―
* gcal_run.ico : アイコンファイル
* install_shortcut.ps1 : ショートカット作成用スクリプト

#### ツールのインストール

解凍したフォルダに移動してください。

installer.exeをダブルクリックしてください。

指示に従ってインストールしてください。

```
------------------------------------------------
GoogleカレンダーTV会議強制起動ツールインストラ―
バージョン: v2.2.3
------------------------------------------------
現状、GoogleカレンダーTV会議強制起動ツールはまだインストールされていません

ツールをインストールしますか？ (y/n) > 
```

yを入力します

```
GoogleカレンダーAPIのクレデンシャルパスを指定してください > 
```

クレデンシャルファイルを指定してください。例えば、デスクトップに「gcal.json」というファイル名で保存した場合は、`C:\Users\ユーザ名\Desktop\gcal.json` となります。

```
ブラウザアプリケーションのパスを指定してください
デフォルトは「C:\Program Files (x86)\Google\Chrome\Application\chrome.exe」です。デフォルトで良い場合は何も入力せずにEnterを押してください
>
```

ツールで使うブラウザを指定してください。デフォルトはGoogle Chromeです。そのままでよければ何も入力せずにEnterを押してください。



```

会議の何分前に起動するか指定してください
デフォルトは「2分」です。デフォルトで良い場合は何も入力せずにEnterを押してください
>
入力された文字列:

インストール先ディレクトリを作成しました: C:\Users\fetaro\AppData\Roaming/gcal_run
設定ファイルを作成しました: C:\Users\fetaro\AppData\Roaming/gcal_run/config.json
ツールをインストールディレクトリにコピーします. "." -> "C:\Users\fetaro\AppData\Roaming/gcal_run"
ブラウザを使ってこのアプリケーションを認証してください。URL = https://accounts.google.com/o/oauth2/auth?client_id=519759873450-hvio180ddihq7832531jl0c9uedf9m19.apps.googleusercontent.com&redirect_uri=http%3A%2F%2F127.0.0.1%3A51080&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fcalendar.readonly&state=st1739259094746220200
[2025-02-11 16:31:43] [ INFO] OAuthトークンをファイルに保存。トークンファイル=C:\Users\fetaro\AppData\Roaming/gcal_run/oauth_token
インストールが完了しました。

プログラムを動かすには %s をダブルクリックして起動してください C:\Users\fetaro\AppData\Roaming/gcal_run/gcal_run

デスクトップにショートカットを作りますか？ (y/n) >



## アンインストール方法

installer.exeをダブルクリックして、アンインストールを選択してください。

## バージョンアップ方法

installer.exeをダブルクリックして、バージョンアップを選択してください。

## インストールに失敗する場合

[よくある質問](https://github.com/fetaro/gcal-run/wiki)を御覧ください




