package route

import (
	"dropout_s_back/config"
	"dropout_s_back/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Init ルーティング
func Init(conn *gorm.DB) *gin.Engine {
	r := gin.Default()

	ctrler := controller.NewController(conn)

	// BasicAuthの設定
	ba := config.GetBAConfig()
	accounts := gin.Accounts{
		ba.User: ba.Pass,
	}
	authorized := r.Group("/", gin.BasicAuth(accounts))

	api := authorized.Group("/api")
	{
		api.GET("/ble/get", ctrler.GetBle)
		api.GET("/message/get", ctrler.GetMessage)
		api.POST("/user/signup", ctrler.SignUp)
		api.POST("/message/post", ctrler.PostMessage)
		api.GET("/user/get", ctrler.GetUsers)
	}

	return r
}
