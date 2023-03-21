package middleware

import (
	"bytes"
	"encoding/base64"
	"gin/common/lib/redis"
	userservice "gin/service/userService"
	"io"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		type authStr struct {
			Uid  string `json:"uid"`
			Auth string `json:"token"`
		}
		var o authStr
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.ShouldBind(&o)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		auth := c.GetHeader("Auth")

		if auth == "" {
			auth = o.Auth
		}

		if auth == "" {
			c.Next()
			return
		}

		var uid int64
		authTrue, err := base64.StdEncoding.DecodeString(auth)
		if err == nil {
			arr := strings.Split(string(authTrue), "|")
			if len(arr) == 2 {
				auth = arr[0]
				uid, _ = strconv.ParseInt(arr[1], 10, 64)
			}
		}
		if uid == 0 {
			uidStr := c.GetHeader("Auth-Uid")
			if uidStr == "" {
				uidStr = o.Uid
			}
			uid, _ = strconv.ParseInt(uidStr, 10, 64)
		}

		key := redis.GetRedisKey("OPERATE-TOKEN", uid)
		var userTokenRedis userservice.UserTokenRedis
		redis.GetObject(key, &userTokenRedis)
		bol, msg := userservice.CheckToken(uid, auth, userTokenRedis)
		if bol {
			c.Set("uid", uid)
			c.Set("token", auth)
		} else {
			c.Set("msg", msg)
			c.Next()
		}

		c.Next()
	}
}
