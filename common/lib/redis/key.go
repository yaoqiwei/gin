package redis

import (
	"context"
	"gin/model/http_error"
	"gin/util/stringify"
	"runtime"
	"time"
)

type RedisKeyCode string

var RedisKey = map[RedisKeyCode]string{
	"OPERATE-TOKEN": "operate:token",
	"STATISTICS":    "statistics",
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

func tryLock(keyCode RedisKeyCode, others ...interface{}) (ok bool, err error) {
	set, err := SetNX(GetRedisKey(keyCode, others...), 1, 5*time.Second)
	return set, err
}

func SpinLock(ctx context.Context, keyCode RedisKeyCode, others ...interface{}) (timeOut bool, err error) {
	var (
		ok bool
	)
	for {
		select {
		case <-ctx.Done():
			timeOut = true
			return

		default:
			ok, err = tryLock(keyCode, others...)
			if err != nil {
				return
			}

			if ok {
				return
			} else {
				runtime.Gosched()
			}
		}
	}
}
