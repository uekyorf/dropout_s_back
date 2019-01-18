package controller

import (
	"github.com/gin-gonic/gin"
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

// GetBle データベースにあるBLEの一覧を返す
func (ctrler Controller) GetBle(c *gin.Context) {
	//db := ctrler.conn //DB接続
}
