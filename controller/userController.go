package controller

import (
	"dropout_s_back/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestSignUp ユーザ登録apiへのリクエスト
type RequestSignUp struct {
	Name   string `json:"user_name"`
	Device string `json:"device_name"`
}

// SignUp 要求に基づいてユーザを作成する
func (ctrler Controller) SignUp(c *gin.Context) {
	dbConn := ctrler.conn //DB接続
	req := RequestSignUp{}
	err := c.BindJSON(&req)

	// requestがjsonとして正しい構造であるか否か
	if err != nil {
		response := CreateResponse(400, "bad request", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	// requestが条件を満たしているか否か
	if req.Device == "" || req.Name == "" {
		response := CreateResponse(400, "bad request", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	// 要求されたユーザは重複していないか
	var existUsers []db.User
	count := 0
	dbConn.Where("name=?", req.Name).Find(&existUsers).Count(&count)
	if count != 0 {
		response := CreateResponse(409, "user already exist", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	// 登録処理
	newUser := db.User{Name: req.Name}
	dbConn.Create(&newUser)
	newDevice := db.Device{Name: req.Device, UserID: newUser.ID}
	dbConn.Create(&newDevice)

	response := CreateResponse(200, "user created", nil)
	c.JSON(http.StatusOK, response)
}

// GetUsers ユーザ検索API
func (ctrler Controller) GetUsers(c *gin.Context) {

}
