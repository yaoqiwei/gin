package middleware

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {

	//TODO 缺少redis比较
	return func(c *gin.Context) {
		type authStr struct {
			Uid  string `json:"uid`
			Auth string `json:"token"`
		}
		var o authStr
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.ShouldBind(&o)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
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
		fmt.Println(uid)
		// if uid != 0 {
		// 	msg, bol = userService.CheckMenuAuth(c.Request.URL.Path, uid, c, userTokenRedis)
		// 	if !bol {
		// 		c.Set("msg", msg)
		// 	}
		// }
		c.Next()
	}
}
