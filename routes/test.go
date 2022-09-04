package routes

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func Test(apiGroup *gin.RouterGroup) {
	apiUserGroup := apiGroup.Group("/test")
	// 用户登入
	apiUserGroup.POST("/login", controller.TestGet)
	apiUserGroup.POST("/cron", controller.CornTest)
	apiUserGroup.POST("/order", controller.OrderNum)

}
