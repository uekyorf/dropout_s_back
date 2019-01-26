package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// Init データベースの接続をする。
func Init() {

	db, err = gorm.Open("mysql", "b56e5ef47c5c33:c07f2895@tcp(us-cdbr-iron-east-03.cleardb.net)/heroku_69b0c80d331b1b0?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	db.Set("gorm:table_option", "ENGINE=InnoDB")
	db.AutoMigrate(&Device{}, &User{}, &Ble{}, &Message{}, &SendMessage{})

}

// GetConn DBのコネクションを返す。
func GetConn() *gorm.DB {
	db.DB().SetMaxIdleConns(0)
	return db
}
