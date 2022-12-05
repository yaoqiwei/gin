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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (*TestController) SendTest(c *gin.Context) {
	var p body.SendParam
	request.Bind(c, &p)
	aa := isAnagram("anagrama", "nagarama")
	middleware.Success(c, aa)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charMap := make([]int, 26)
	for _, v := range s {
		charMap[v-'a']++
	}
	for _, v := range t {
		if charMap[v-'a']--; charMap[v-'a'] < 0 {
			return false
		}
	}
	return true
}
