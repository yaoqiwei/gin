package request

import (
	"encoding/json"
	"gin/model/http_error"
	"io"

	"github.com/gin-gonic/gin"
)

// Bind 绑定传入数据到结构体
func Bind(c *gin.Context, i interface{}) {
	err := c.ShouldBind(i)
	if err != nil && err != io.EOF {

		terr := http_error.MissingParametersError
		if e, ok := err.(*json.UnmarshalTypeError); ok {
			terr.ErrorMsg += ", 参数类型错误: " + e.Field
		}
		panic(terr)
	}
}
