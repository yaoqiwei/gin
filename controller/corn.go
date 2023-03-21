package controller

import (
	"fmt"
	"gin/common/lib"
	"gin/middleware"
	"gin/model/body"
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
	// pushrecord.PushrecordAdd(p)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copy(a[3:], a[4:])
	fmt.Println(a)
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
	result := strStr("bbaaaccc", "aaac")
	fmt.Println("result", result)
	middleware.Success(c)
}
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}
