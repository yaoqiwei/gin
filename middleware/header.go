package middleware

import (
	"bytes"
	"gin/config"
	"gin/constant"
	"gin/util/cryp"
	"io/ioutil"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func HeaderAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.Http.HeaderCheck {
			c.Next()
			return
		}
		timestamp := c.GetHeader("timestamp")

		if timestamp == "" {
			panic("block_4")
		}

		ti, _ := strconv.ParseInt(timestamp, 10, 64)
		now := time.Now().Unix()

		if ti-10 > now || ti+10 < now {
			panic("block_3")
		}

		type endData struct {
			EndData string `json:"endData"`
		}

		var o endData
		c.ShouldBind(&o)
		if o.EndData == "" {
			panic("block_1")
		}

		str, _ := cryp.AesEncrypt(o.EndData, constant.ApiAesKey)

		if str == "" {
			panic("block_2")
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(str)))
		c.Next()
	}

}
