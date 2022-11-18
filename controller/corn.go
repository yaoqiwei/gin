package controller

import (
	"fmt"
	"gin/common/lib"
	"gin/middleware"
	"gin/model/body"
	"gin/service/pushrecord"
	"gin/util/request"
	"regexp"

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
	str := "Golang reguest expressions for testing"
	matched, err := regexp.MatchString("^Golang", str)
	fmt.Println("matched", matched, err)
}
