package controller

import (
	"github.com/jinzhu/gorm"
)

// Controller Controller全体で使いたいものを定義する
type Controller struct {
	conn *gorm.DB
}

// NewController Controllerにいろいろ入れて返す
func NewController(conn *gorm.DB) Controller {
	return Controller{conn: conn}
}
<<<<<<< HEAD
=======

// GetBle データベースにあるBLEの一覧を返す
func (ctrler Controller) GetBle(c *gin.Context) {
	//db := ctrler.conn //DB接続
}

// CreateResponse APIのレスポンスを生成してくれる。
func CreateResponse(code int, message string, result interface{}) gin.H {
	return gin.H{"code": code, "message": message, "result": result}
}
>>>>>>> 9332026fff60054806dfbbfc5f493ae1d5af082a
