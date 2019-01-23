# dropout_s_back
* Hack U 参加作品のバックエンド/サーバサイド
## 環境
* Golang
  * GORM
  * Gin
  * mysql
  
## 実行
### データベースの設定
* 以下のファイルにDB接続に必要な情報を記述
```
/config/dbconfig.go
```

### Basic認証の設定
* 以下のファイルのユーザとパスワードを編集
```
/config/baconfig.go
```

### パッケージの入手
```
go get "github.com/go-sql-driver/mysql"
go get "github.com/jinzhu/gorm"
go get "github.com/gin-gonic/gin"
```

### run
```
go run main.go
```

## APIリスト
### ユーザ登録
* https://github.com/EIMIKI/dropout_s_back/pull/20
### ユーザ検索
* https://github.com/EIMIKI/dropout_s_back/pull/25
### メッセージ投稿
* https://github.com/EIMIKI/dropout_s_back/pull/22
### メッセージ取得
* coming soon
### BLE取得
* https://github.com/EIMIKI/dropout_s_back/pull/15
