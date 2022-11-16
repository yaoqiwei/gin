package middleware

import (
	"bytes"
	"gin/util/jwt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path
		start := time.Now()
		method := c.Request.Method
		var body string

		contentType := c.Request.Header.Get("content-type")
		for _, v := range strings.Split(contentType, ";") {
			if v == "application/json" {
				bodyBytes, _ := io.ReadAll(c.Request.Body)

				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

				body = string(bodyBytes)
				if body != "{}" && body != "" {
					body = ", body:" + body
				} else {
					body = ""
				}
			}
		}
		defer func() {
			var user string
			uid := jwt.GetUid(c)
			uidStr := strconv.FormatInt(uid, 10)
			if uid > 0 {
				user = c.ClientIP() + ", " + uidStr
			} else {
				user = c.ClientIP()
			}

			logrus.Infof("%s %s:%s%s", user, method, path, body)

			latency := time.Since(start)
			if latency > time.Second {
				logrus.Warnf("%s %s:%s, latency:%v", user, method, path, latency)
			}
		}()

		c.Next()

	}
}
