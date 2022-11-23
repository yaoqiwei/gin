package controller

import (
	"container/list"
	"fmt"
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
	treeNode := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	aa := maxDepth(&treeNode)
	middleware.Success(c, aa)
}

func PrintlnList(l *list.List) {
	if l.Front() == nil {
		return
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func maxDepth(root *TreeNode) int {
	// if root == nil {
	// 	return 0
	// }
	// left := maxDepth(root.Left)
	// right := maxDepth(root.Right)
	// max := math.Max(float64(left), float64(right))
	// return int(max) + 1

	var result int
	var quene = list.New()
	if root == nil {
		return 0
	}
	quene.PushBack(root)
	for quene.Len() > 0 {

		length := quene.Len()
		for i := 0; i < length; i++ {
			node := quene.Remove(quene.Front()).(*TreeNode)
			if node.Left != nil {
				quene.PushBack(node.Left)
			}
			if node.Right != nil {
				quene.PushBack(node.Right)
			}
		}
		result += 1
	}
	return result

}
