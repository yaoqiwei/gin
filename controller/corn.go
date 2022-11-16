package controller

import (
	"gin/common/lib"
	"gin/middleware"
	"gin/model/body"
	"gin/service/pushrecord"
	"gin/util/request"

	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func TestRegister(router lib.RegisterRouterGroup, needLoginRouter lib.RegisterRouterGroup) {
	c := TestController{}
	needLoginRouter.POST("/test", c.CornTest)
	needLoginRouter.POST("/send", c.SendTest)
}

func (*TestController) CornTest(c *gin.Context) {
	var p body.PushrecordParam
	request.Bind(c, &p)
	pushrecord.PushrecordAdd(p)
	middleware.Success(c)
}

func (*TestController) SendTest(c *gin.Context) {

}
