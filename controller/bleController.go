package controller

import (
	"dropout_s_back/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseBles 返却するBLE情報
type ResponseBles struct {
	Name     string `json:"ble_uuid"`
	AreaName string `json:"area_name"`
}

// GetBleAll データベースにあるBLEの一覧を返す
func (ctrler Controller) GetBleAll(c *gin.Context) {
	dbConn := ctrler.conn //DB接続
	ctrler.mux.Lock()
	defer ctrler.mux.Unlock()
	c.Header("Access-Control-Allow-Origin", "*")

	result := []ResponseBles{}
	var bles []db.Ble
	var count int

	// BLEリスト作成
	dbConn.Find(&bles).Count(&count)
	for _, ble := range bles {
		name := ble.Name
		areaName := ble.AreaName
		responseBle := ResponseBles{Name: name, AreaName: areaName}

		result = append(result, responseBle)
	}

	// BLEが一つも登録されていない時
	if count == 0 {
		response := CreateResponse(404, "ble not found", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := CreateResponse(200, "request completed", result)
	c.JSON(http.StatusOK, response)
}
