package middleware

import (
	"gin/util/compute"
	"gin/util/convert"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	SuccessCode ResponseCode = iota
	ErrorCode
	EncryptionErrorCode ResponseCode = -99
)

type Response struct {
	ErrorCode ResponseCode `json:"code"`
	ErrorMsg  string       `json:"msg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"traceId"`
	Stack     interface{}  `josn:"stack,omitempty"`
}

// Error 错误处理
func Error(c *gin.Context, code ResponseCode, msg string) {
	resp := &Response{ErrorCode: code, ErrorMsg: msg}
	SerializeJson(c, resp, 200)
	//Abort 在被调用的函数中阻止挂起函数。之后的函数不会执行
	c.Abort()
}

// Success 正常返回
func Success(c *gin.Context, content ...interface{}) {
	var data interface{}
	if len(content) > 0 {
		data = content[0]
	}
	resp := &Response{ErrorCode: SuccessCode, ErrorMsg: "", Data: data}
	SerializeJson(c, resp, 200)
}

// SerializeJson 序列化json
func SerializeJson(c *gin.Context, res *Response, httpCode int) {
	res.TraceId = compute.GetRandomString(6)
	if res.Data == nil {
		res.Data = struct{}{}
	}
	c.Set("response", convert.ToJson(res))
	c.JSON(httpCode, res)
}
