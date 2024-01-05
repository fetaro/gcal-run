Googleカレンダーイベント自動起動ツール(gcal_run)
======================

## これは何？

GoogleカレンダーにあるWeb会議のイベントを自動的にブラウザで起動するツールです。 Web会議に遅れがちな人にオススメです。

インストールするとMacOSのバックグラウンドで動作します。MacOSのUserAgentに登録されるので、Macを再起動しても自動的に起動します。

カレンダーを定期的にチェックし、GoogleMeet, Zoom, TeamsのURLがイベントにある場合は、イベント開始時間の2分前にブラウザで自動的に起動します。


## インストール方法

[このマニュアル](https://developers.google.com/calendar/api/guides/overview?hl=ja)に従いGoogle Calendar APIのクレデンシャルファイルを取得してください。
クレデンシャルファイルはJSON形式です。

以下のコマンドでインストーラを実行してください。
```text
$ ./bin/installer (クレデンシャルファイルのパス)
```

プロンプトの指示に従ってインストールしてください。

インストールが成功すると、指定したインストールディレクトリに実行ファイル等がインストールされるとともに、 MacのUserAgentに `/Users/fetaro/Library/LaunchAgents/com.github.fetaro.gcal_run.plist` として登録されます。

## 使い方

以下のコマンドでバックグラウンドプロセスを起動してください。
```text
$ launchctl load /Users/fetaro/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

バックグラウンドプロセスを終了する場合は以下のコマンドを実行してください。
```text
$ launchctl unload /Users/fetaro/Library/LaunchAgents/com.github.fetaro.gcal_run.plist
```

ログは `(インストールディレクトリ)/gcal_run.log` に出力されます。

アンインストールする場合はデーモンを終了した後、インストールディレクトリを削除してください。




