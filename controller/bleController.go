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

// GetBle データベースにあるBLEの一覧を返す
func (ctrler Controller) GetBle(c *gin.Context) {
	dbConn := ctrler.conn //DB接続

	response := []ResponseBles{}
	var bles []db.Ble

	dbConn.Find(&bles)
	for _, ble := range bles {
		name := ble.Name
		areaName := ble.AreaName
		responseBle := ResponseBles{Name: name, AreaName: areaName}

		response = append(response, responseBle)
	}

	c.JSON(http.StatusOK, response)
}
