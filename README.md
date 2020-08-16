# backend-summer-vacation
バックエンドの夏休み間に使う練習リポジトリ
## 実行方法
- ローカル実行コマンド：`make`
- コンテナイメージビルドコマンド：`make docker-build`
- コンテナ実行コマンド：`make docker-run`
- コンテナ削除コマンド：`make docker-clean`

## 課題内容
### チュートリアル
**0-1 Hello World**  
http://localhost:8080 にリクエストを投げて以下のJSONが返ってくることを確認する  
レスポンス内容
```json
{
    "message": "hello world"
}
```
**0-2 Say Hello**  
http://localhost:8080/sayhello にリクエストを投げて以下のJSONが返ってくることを確認する  
リクエスト内容  
```json
{
    "name": "hoge"
}
```
レスポンス内容
```json
{
    "message": "hello hoge"
}
```
### 課題1
現在の日付と時間を返す処理を作成  
**リクエスト**  
HTTPメソッド： GET  
**レスポンス**  
```json
{
  "timestamp": "string",
  "detail": {
    "date": "string",
    "time": "string"
  }
}
```
例: date → 2020-09-02
例: time → 00:00:00
### 課題2
ツェラーの公式でリクエストで投げた日付の曜日を返す  
**リクエスト**  
HTTPメソッド： POST  
```json
{
  "year": "Int",
  "month": "Int",
  "day": "Int",
}
```
**レスポンス**  
```json
{
  "week": "string"
}
```
例： week → Monday
### 課題3
ユーザーIDとパスワードをデータベースに登録して, 発行したトークンを返す  
パスワードはハッシュ化したものをデータベースに登録する  
**リクエスト**  
HTTPメソッド： POST  
```json
{
  "id": "string",
  "password": "string",
}
```
**レスポンス**  
```json
{
  "token": "string"
}
```
### 課題4
ユーザーIDとパスワードをデータベースに登録されたものかを照合する
照合が終わったら結果を返す  
**リクエスト**  
HTTPメソッド： POST  
```json
{
  "id": "string",
  "password": "string"
}
```
**レスポンス**  
```json
{
  "certification": "boolean"
}
```