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
