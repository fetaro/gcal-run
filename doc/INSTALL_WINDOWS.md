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

#### ツールのインストール

解凍したフォルダに移動してください。

installer.exeをダブルクリックしてください。

指示に従ってインストールしてください。

以下の設問が表示された、上記で入手したクレデンシャルファイルを指定してください。

```
GoogleカレンダーAPIのクレデンシャルパスを指定してください > 
```

例えば、デスクトップに「gcal.json」というファイル名で保存した場合は、`C:\Users\ユーザ名\Desktop\gcal.json` となります。

以下の設問が表示されたら、利用するブラウザのパスを指定してください。

```
ブラウザアプリケーションのパスを指定してください
デフォルトは「C:\Program Files (x86)\Google\Chrome\Application\chrome.exe」です。デフォルトで良い場合は何も入力せずにEnterを押してください
>
```

デフォルトはGoogle Chromeです。そのままでよければ何も入力せずにEnterを押してください。

Windowsに最初から入っているブラウザ「Edge」を使いたい場合は、 `C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe` と入力してください。

ブラウザが起動して、Googleの認証画面が表示されます。TV会議を強制起動したいGoogleカレンダーのアカウントでログインし、
以下のようなアクセスを求める画面で「続行」を選択してください。

![image](2.png)

最後に、以下の質問で、自動で起動されるようにスタートアップに登録するか聞かれるため、
自動で希望する場合はyを入力してください。

```
自動で起動されるように、スタートアップに登録しますか？ (y/n) > y
```

自動で起動をされないようにしたい場合は、nを入力してください。

その場合、ツールを自動で起動する必要がありますが、その場合はデスクトップにある「gcal_run」のショートカットをダブルクリックして起動してください


## アンインストール方法

installer.exeをダブルクリックして、アンインストールを選択してください。

## バージョンアップ方法

installer.exeをダブルクリックして、バージョンアップを選択してください。

## インストールに失敗する場合

[よくある質問](https://github.com/fetaro/gcal-run/wiki)を御覧ください




