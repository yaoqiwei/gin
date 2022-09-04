package controller

import (
	"gin/context"
	"gin/model/body"
	"gin/service/order"

	"github.com/gin-gonic/gin"
)

func OrderNum(c *gin.Context) {
	var p body.OrderNumParam
	if err := c.ShouldBind(&p); err != nil {
		context.ValidateError(c)
		return
	}
	order.OrderSuccess(p)
}
