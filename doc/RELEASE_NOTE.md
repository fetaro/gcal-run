リリースノート
======================
### v2.3
* オフラインになってもツールが停止しないようにした。これによりWindowsでスリープからの復旧時にツールが停止しないようになった。
* インストラーでクレデンシャルファイルやブラウザを選択しやすくした。

### v2.1
* Windows版のリリース
* インストーラーを改善

### v2.0
* 30秒ごとにカレンダーをチェックするようにした。
* アップローダーと、アンインストラーの追加。
* インストールディレクトリを変更できないようにした。インストールディレクトリは `${HOME}/.gcal_run/` に固定される。
* LaunchAgentsの設定ファイルを変更。コマンドのオプションを指定しないようにした。代わりにコマンドのオプションは `${HOME}/.gcal_run/config.json` に記述するようにした。
* ログの改善。

### v1.1 
* 初期バージョン 
