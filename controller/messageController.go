package controller

import (
	"github.com/gin-gonic/gin"
)

// GetMessage 要求(User,BLE)に基づいてメッセージを返却する
func (ctrler Controller) GetMessage(c *gin.Context) {
	//db := ctrler.conn //DB接続

}

// PostRequestの構造体
type PostRequestJson struct {
	Device_name string   `json:"device_name"`
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	Due         string   `json:"due"`
	Ble_uuid    string   `json:"ble_uuid"`
	To_user     []string `json:"to_user"`
}

// PostMessage 要求に基づいてメッセージをデータベースに登録する
func (ctrler Controller) PostMessage(c *gin.Context) {
	//db := ctrler.conn //DB接続

}
