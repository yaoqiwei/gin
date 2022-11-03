package jwt

import (
	"encoding/base64"
	"gin/model/http_error"
	"gin/util/cryp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Token struct {
	Token       string
	EncodeToken string
	Uid         int64
	UserLogin   string
	Expire      time.Time
}

func SetPass(pass string, authcode string) string {
	return cryp.MD5(authcode + pass)
}

func SetToken(uid int64, userLogin string) *Token {

	t := time.Now()
	token := cryp.MD5(cryp.MD5(strconv.FormatInt(uid, 10) + userLogin + strconv.FormatInt(t.Unix(), 10)))

	encodeToken := base64.StdEncoding.EncodeToString([]byte(token + "|" + strconv.FormatInt(uid, 10)))

	return &Token{
		Token:       token,
		EncodeToken: encodeToken,
		Uid:         uid,
		UserLogin:   userLogin,
		Expire:      t,
	}
}

func GetUid(c *gin.Context, e ...bool) int64 {
	var i int64
	if uid, exists := c.Get("uid"); exists {
		i = uid.(int64)
	}
	if len(e) > 0 && e[0] {
		i = int64(64)
		//if i == 0 {
		//	panic(http_error.JwtError)
		//}
		//if c.GetString("msg") != "" {
		//	var ret = http_error.HttpError{
		//		ErrorCode: 701,
		//		ErrorMsg:  c.GetString("msg"),
		//	}
		//	panic(ret)
		//}
		//i = 53
	}
	return i
}

func GetToken(c *gin.Context, e ...bool) string {
	var i string
	if token, exists := c.Get("token"); exists {
		i = token.(string)
	} else if len(e) > 0 && e[0] {
		panic(http_error.JwtError)
	}
	return i
}
