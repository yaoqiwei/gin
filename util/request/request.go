package request

import (
	"encoding/json"
	"gin/model/http_error"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Bind 绑定传入数据到结构体
func Bind(c *gin.Context, i interface{}) {
	err := c.ShouldBind(i)
	if err != nil && err != io.EOF {

		terr := http_error.MissingParametersError
		if e, ok := err.(*json.UnmarshalTypeError); ok {
			terr.ErrorMsg += ", 参数类型错误: " + e.Field
		}

		if e, ok := err.(validator.ValidationErrors); ok {
			terr.ErrorMsg += ",参数校验错误" + e.Error()
		}
		panic(terr)
	}
}
