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

// CreateResponse APIのレスポンスを生成してくれる。
func CreateResponse(code int, message string, result interface{}) gin.H {
	return gin.H{"code": code, "message": message, "result": result}
}
