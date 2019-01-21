package controller

import (
	"dropout_s_back/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetMessage 要求(User,BLE)に基づいてメッセージを返却する
func (ctrler Controller) GetMessage(c *gin.Context) {
	//db := ctrler.dbdbConn //DB接続

}

// PostRequestの構造体
type RequestMessagePost struct {
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
	req := RequestMessagePost{}
	err := c.BindJSON(&req)

	// requestがjsonとして正しい構造であるか否か
	if err != nil {
		response := CreateResponse(400, "bad request", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	// requestが条件を満たしているか否か
	if req.Device_name == "" || req.Title == "" || req.Body == "" || req.Ble_uuid == "" || len(req.To_user) == 0 {
		response := CreateResponse(400, "bad request", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	// リクエストの内容を基にSELECT
	device := db.Device{}
	if dbConn.Where("name=?", req.Device_name).First(&device).RecordNotFound() {
		response := CreateResponse(404, "device is not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	user := db.User{}
	if dbConn.First(&user, device.UserID).RecordNotFound() {
		response := CreateResponse(404, "A user using the device is not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	ble := db.Ble{}
	if dbConn.Where("name=?", req.Ble_uuid).First(&ble).RecordNotFound() {
		response := CreateResponse(404, "BLE is not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	// messageをINSERT
	message := db.Message{}
	message.UserID = user.ID
	message.Title = req.Title
	message.Body = req.Body
	message.BleID = ble.ID
	if req.Due == "" {
		t := time.Now().AddDate(0, 1, 0)
		req.Due = t.Format("2006-01-02-15-04")
	}
	message.Due, _ = time.Parse("2006-01-02-15-04-05 MST", req.Due+"-00 JST")
	dbConn.Create(&message)

	// sendMessageをINSERT
	for _, value := range req.To_user {
		toUser := db.User{}
		if dbConn.Where("name=?", value).First(&toUser).RecordNotFound() {
			response := CreateResponse(404, "Recipient is not found", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		sendMessage := db.SendMessage{}
		sendMessage.MessageID = message.ID
		sendMessage.UserID = toUser.ID
		dbConn.Create(&sendMessage)
	}
	response := CreateResponse(200, "Submitted message", nil)
	c.JSON(http.StatusOK, response)
}
