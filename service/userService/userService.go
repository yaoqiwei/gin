package userservice

import (
	"gin/common/lib/redis"
	"gin/dao/mapper/users"
	"gin/model/body"
	"gin/model/http_error"
	"time"
)

// UserTokenRedis 用户token redis存储结构体
type UserTokenRedis struct {
	Uid        int64  `json:"uid"`        // 用户id
	Token      string `json:"token"`      // 用户token
	UserType   byte   `json:"userType"`   // 0 普通用户，1 admin
	Expiretime int64  `json:"expiretime"` // token到期时间
}

// CheckToken 校验token
func CheckToken(uid int64, token string, userTokenRedis UserTokenRedis) (bool, string) {
	if uid == 0 || token == "" {
		return false, ""
	}
	if userTokenRedis.Uid == 0 {
		user, _ := users.Get(body.UserSearchParam{Uid: uid})
		if user.Id != 0 {
			userTokenRedis.Uid = user.Id
			userTokenRedis.Token = user.Token
			userTokenRedis.Expiretime = user.ExpireTime
			userTokenRedis.UserType = user.UserType
			redis.Set(redis.GetRedisKey("OPERATE-TOKEN", uid), userTokenRedis, 60*60*24*300)
		}
	}
	if userTokenRedis.Uid != 0 {
		// 账号已经在其它设备登录
		if userTokenRedis.Token != token {
			return false, http_error.JwtError.ErrorMsg
		}
		// 你的登录状态过期，请重新登录
		if userTokenRedis.Expiretime < time.Now().Unix() {
			return false, http_error.JwtError.ErrorMsg
		}

		return true, ""
	}
	return false, ""
}
