package controller

import (
	"dropout_s_back/db"
	"time"

	"github.com/gin-gonic/gin"
)

// GetMessage 要求(User,BLE)に基づいてメッセージを返却する
func (ctrler Controller) GetMessage(c *gin.Context) {
	//db := ctrler.dbdbConn //DB接続

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
	//DB接続
	dbConn := ctrler.conn

	// リクエストをバインド
	req := PostRequestJson{}
	c.BindJSON(&req)
	// リクエストの内容を基にSELECT
	device := db.Device{}
	dbConn.Where("name=?", req.Device_name).First(&device)
	user := db.User{}
	dbConn.First(&user, device.UserID)
	ble := db.Ble{}
	dbConn.Where("name=?", req.Ble_uuid).First(&ble)
	// messageを作成し、INSERT
	message := db.Message{}
	message.UserID = user.ID
	message.Title = req.Title
	message.Body = req.Body
	message.BleID = ble.ID
	if req.Due == "" {
		t := time.Now()
		req.Due = t.Format("2006-01-02-15-04")
	}
	message.Due, _ = time.Parse("2006-01-02-15-04-05 MST", req.Due+"-00 JST")
	message.Due = message.Due.AddDate(0, 1, 0)
	dbConn.Create(&message)

}
