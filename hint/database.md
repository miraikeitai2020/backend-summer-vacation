# データベースに接続してみよう
## GoでMySQLに接続する
いい感じにMySQLに接続できるライブラリがある
- [database/sql](https://golang.org/pkg/database/sql/)（Goの標準ライブラリ）
- [github.com/jinzhu/gorm](https://godoc.org/github.com/jinzhu/gorm)

好きなの選んでいいよ
## DBに接続するためのヒント
### データベース情報
| | |
|:-:|:-|
|rootパスワード|miraiketai2020_admin|
|ユーザアカウント|miraiketai2020|
|ユーザパスワード|miraiketai2020|
|データベース|summer|
|IP|db|
|ポート|3306|

接続情報が書かれているファイルは以下の2つ
- `docker-compose.yml`
- `Dockerfile`

接続情報を取得するためのプログラム
- `pkg/config/config.go`

### データベースとAPIを起動する
dockerで用意しました  
以下，実行コマンド  
makeコマンドを使用する場合  
- サービス実行コマンド: `make service-build`
- サービス停止コマンド: `make service-clean`

makeコマンドを使用しない場合  
- サービス構築コマンド: `docker-compose build`
- サービス実行コマンド: `docker-compose up -d` 
- サービス停止コマンド: `docker-compose stop`

### 接続に必要なもの
- 接続に必要な情報（トークン）
`config.GetConnectionToken()`で取得できます
