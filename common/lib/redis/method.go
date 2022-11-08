package redis

import (
	"time"

	"github.com/sirupsen/logrus"
)

// GetString 获取string类型键值
func GetString(key, clusterName string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

// GetInt 获取int类型的键值
func GetInt(key, clusterName string) (int, error) {
	return RDB.Get(ctx, key).Int()
}

// GetInt64 获取int64类型键值
func GetInt64(key, clusterName string) (int64, error) {
	return RDB.Get(ctx, key).Int64()
}

// GetBool 获取bool类型键值
func GetBool(key, clusterName string) (bool, error) {
	return RDB.Get(ctx, key).Bool()
}

// Set 对应key值设置键值
func Set(key, value string, params ...time.Duration) error {
	var expiration time.Duration
	if len(params) > 0 {
		expiration = params[0]
	} else {
		expiration = DefaultExpiration
	}
	err := RDB.Set(ctx, key, value, expiration).Err()
	if err != nil {
		logrus.Error("Redis SET", err)
	}
	return err
}

// SetNX 添加key值，如果存在则添加失败，不会修改原来的值
func SetNX(key, value string, params ...time.Duration) (bool, error) {
	var expiration time.Duration
	if len(params) > 0 {
		expiration = params[0]
	} else {
		expiration = DefaultExpiration
	}
	bool, err := RDB.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		logrus.Error("Redis SetNX", err)
	}
	return bool, err
}

// Exists 对应key值是否存在
func Exists(key ...string) (int64, error) {
	return RDB.Exists(ctx, key...).Result()
}

// Del 删除对应key值的数据
func Del(key ...string) error {
	_, err := RDB.Del(ctx, key...).Result()
	return err
}

// Flush 删除所有数据
func Flush() error {
	_, err := RDB.FlushAll(ctx).Result()
	return err
}
