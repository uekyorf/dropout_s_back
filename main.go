package main

import (
	"dropout_s_back/db"
)

func main() {
	db.Init()            // DB接続初期化
	conn := db.GetConn() // DB接続取得
	defer conn.Close()

	r := Init(conn) // routes初期化
	r.Run(":3000")  // サーバ起動
}
