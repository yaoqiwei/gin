package middleware

import (
	"fmt"
	"gin/config"
	"gin/response"
	"gin/response/http_error"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RecoveryMiddleware buhuo
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		defer func() {
			if err := recover(); err != nil {
				latency := time.Since(start)
				if e, ok := err.(http_error.HttpError); ok {
					logrus.Infof("response:[%d,%v] %s", e.ErrorCode, latency, e.ErrorMsg)
					response.Error(c, response.ResponseCode(e.ErrorCode), e.ErrorMsg)
					return
				}
				//先做一下日志记录
				logrus.Warnf("", map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})

				if config.DebugMode != "debug" {
					response.Error(c, 500, "内部错误")
					return
				} else {
					response.Error(c, 500, fmt.Sprint(err))
					return
				}
			}
		}()
		c.Next()
	}
}
