package redis

import (
	"gin/model/http_error"
	"gin/util/stringify"
)

type RedisKeyCode string

var RedisKey = map[RedisKeyCode]string{
	"OPERATE-TOKEN": "operate:token",
}

// GetRedisKey 获取对应key值
func GetRedisKey(keyCode RedisKeyCode, others ...interface{}) string {
	key, ok := RedisKey[keyCode]
	if !ok {
		panic(http_error.NoRedisKey)
	}

	for _, v := range others {
		key += ":" + stringify.ToString(v)
	}
	return key
}

// Lock 加锁
func Lock(keyCode RedisKeyCode, others ...interface{}) {
	set, _ := SetNX(GetRedisKey(keyCode, others...), 1, 5)
	if !set {
		panic(http_error.FrequentOperations)
	}
}

// UnLock 解锁
func UnLock(keyCode RedisKeyCode, others ...interface{}) {
	Del(GetRedisKey(keyCode, others...))
}
