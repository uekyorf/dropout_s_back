package controller

import (
	"dropout_s_back/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestMessageGet struct {
	Ble_uuid  string `form:"ble_uuid"`
	User_name string `form:"user_name"`
}
type ResponseMessageGet struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
}

// GetMessage 要求(User,BLE)に基づいてメッセージを返却する
func (ctrler Controller) GetMessage(c *gin.Context) {
	dbConn := ctrler.conn //DB接続
	req := RequestMessageGet{}
	err := c.ShouldBind(&req)

	// requestが正しい構造であるか
	if err != nil {
		response := CreateResponse(400, "bad request", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	// requestが条件を満たしているか
	if req.Ble_uuid == "" || req.User_name == "" {
		response := CreateResponse(400, "bad request", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	// リクエストの内容を基にSELECT
	ble := db.Ble{}
	if dbConn.Where("name=?", req.Ble_uuid).First(&ble).RecordNotFound() {
		response := CreateResponse(404, "BLE is not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	user := db.User{}
	if dbConn.Where("name=?", req.User_name).First(&user).RecordNotFound() {
		response := CreateResponse(404, "Your name is not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	// send_messagesテーブルからuser_idが一致するレコードをSELECT
	sendMessages := []db.SendMessage{}
	if dbConn.Where("user_id=?", user.ID).Find(&sendMessages).RecordNotFound() {
		response := CreateResponse(404, "Messages to you are not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	// sendMessagesのMessageIDの配列を作成
	// messageテーブルからSELECTするのに使う
	sendMessagesMessageIDs := []uint{}
	for _, sendMessage := range sendMessages {
		sendMessagesMessageIDs = append(sendMessagesMessageIDs, sendMessage.MessageID)
	}
	// messagesテーブルからsendMessagesMessageIDs, ble.IDの一致するレコードをSELECT
	messages := []db.Message{}
	if dbConn.Where("ble_id=?", ble.ID).Find(&messages, sendMessagesMessageIDs).RecordNotFound() {
		response := CreateResponse(404, "Message to you with the BLE is not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}

}

// PostRequestの構造体
type RequestMessagePost struct {
	Device_name  string   `json:"device_name"`
	Title        string   `json:"title"`
	Body         string   `json:"body"`
	Due          string   `json:"due"`
	Ble_uuid     string   `json:"ble_uuid"`
	To_user      []string `json:"to_user"`
	To_all_users bool     `json:"to_all_users" binding:"exists"`
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
	if req.Device_name == "" || req.Title == "" || req.Body == "" || req.Ble_uuid == "" || (len(req.To_user) == 0 && req.To_all_users == false) {
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
	toUser := []db.User{}
	if req.To_all_users == true {
		dbConn.Find(&toUser)
	} else {
		if dbConn.Where("name in (?)", req.To_user).Find(&toUser).RecordNotFound() {
			response := CreateResponse(404, "to user is not found", nil)
			c.JSON(http.StatusOK, response)
			return
		}
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
	message.Due, err = time.Parse("2006-01-02-15-04-05 MST", req.Due+"-00 JST")
	if err != nil {
		response := CreateResponse(400, "bad request", nil)
		c.JSON(http.StatusOK, response)
		return
	}
	dbConn.Create(&message)

	// sendMessageをINSERT
	for key, _ := range toUser {
		sendMessage := db.SendMessage{}
		sendMessage.MessageID = message.ID
		sendMessage.UserID = toUser[key].ID
		dbConn.Create(&sendMessage)
	}
	response := CreateResponse(200, "Submitted message", nil)
	c.JSON(http.StatusOK, response)
}
