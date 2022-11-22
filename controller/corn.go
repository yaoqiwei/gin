package controller

import (
	"gin/common/lib"
	"gin/middleware"
	"gin/model/body"
	"gin/service/pushrecord"
	"gin/util/request"
	"math"

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
	var p body.TestParam
	request.Bind(c, &p)
	aa := reverse(p.Number)
	middleware.Success(c, aa)
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev
}
