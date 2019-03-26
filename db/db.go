package db

import (
	"dropout_s_back/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// Init データベースの接続をする。
func Init() {
	dbConfig := config.GetDBConfig()
	dbUser := dbConfig.User
	dbPass := dbConfig.Pass
	dbName := dbConfig.DBName

	db, err = gorm.Open("mysql", dbUser+":"+dbPass+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
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
