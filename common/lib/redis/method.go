package redis

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// GetString 获取string类型键值
func GetString(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

// GetInt 获取int类型的键值
func GetInt(key string) (int, error) {
	return RDB.Get(ctx, key).Int()
}

// GetInt64 获取int64类型键值
func GetInt64(key string) (int64, error) {
	return RDB.Get(ctx, key).Int64()
}

// GetBool 获取bool类型键值
func GetBool(key string) (bool, error) {
	return RDB.Get(ctx, key).Bool()
}

// GetObject 获取结构体struct的键值
func GetObject(key string, val interface{}) error {
	reply, err := GetString(key)
	if err != nil {
		return err
	}
	return conCode(reply, val)
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

// TTL 当 key 不存在时，返回 -2 。 当 key 存在但没有设置剩余生存时间时，返回 -1 。 否则，以秒为单位，返回 key 的剩余生存时间。
func TTL(key string) (int64, error) {
	time, err := RDB.TTL(ctx, key).Result()
	return int64(time), err
}

// Expire 用于设置 key 的过期时间，key 过期后将不再可用。单位以秒计。
func Expire(key string, expire int64) error {
	_, err := RDB.Expire(ctx, key, time.Duration(expire)).Result()
	return err
}

// Incr 在key中储存的数字+1
func Incr(key string) (int64, error) {
	return RDB.Incr(ctx, key).Result()
}

// IncrBy 将key所储存的值加上给定的增量
func IncrBy(key string, amount int64) (int64, error) {
	return RDB.IncrBy(ctx, key, amount).Result()
}

// Decr 在key中储存的数字-1
func Decr(key string) (int64, error) {
	return RDB.Decr(ctx, key).Result()
}

// DecrBy 在key所存储的值减去给定的值
func DecrBy(key string, amount int64) (int64, error) {
	return RDB.DecrBy(ctx, key, amount).Result()
}

// HMSet 设置用于同时将多个 field-value (字段-值)对设置到哈希表中。此命令会覆盖哈希表中已存在的字段。如果哈希表不存在，会创建一个空哈希表，并执行 HMSET 操作。
func HMSet(key string, val interface{}, expire int64) error {
	pipe := RDB.TxPipeline()
	pipe.HMSet(ctx, key, val)
	pipe.Expire(ctx, key, time.Duration(expire))
	_, err := pipe.Exec(ctx)
	if err != nil {
		errorHandle(err)
		return err
	}
	return nil
}

// Hset 用于为哈希表中的字段赋值 。如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。如果字段已经存在于哈希表中，旧值将被覆盖。
func Hset(key string, val interface{}) error {
	_, err := RDB.HSet(ctx, key, val).Result()
	return err
}

// HsetExpire 设置哈希表中字段赋值，并设置到期时间
func HsetExpire(key string, val interface{}, expire int64) error {
	pipe := RDB.TxPipeline()
	pipe.HSet(ctx, key, val)
	pipe.Expire(ctx, key, time.Duration(expire))
	_, err := pipe.Exec(ctx)
	if err != nil {
		errorHandle(err)
		return err
	}
	return nil
}

// HDel 用于删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略。
func HDel(key string, field ...string) error {
	_, err := RDB.HDel(ctx, key, field...).Result()
	return err
}

// HGetString 用于返回哈希表中指定字段的值(string类型)
func HGetString(key, field string) (string, error) {
	return RDB.HGet(ctx, key, field).Result()
}

// HGetInt HGet的工具方法，当字段值为int类型时使用
func HGetInt(key, field string) (int, error) {
	return RDB.HGet(ctx, key, field).Int()
}

// HGetInt64 HGet的工具方法，当字段值为int64类型时使用
func HGetInt64(key, field string) (int64, error) {
	return RDB.HGet(ctx, key, field).Int64()
}

// HGetBool HGet的工具方法，当字段值为bool类型时使用
func HGetBool(key, field string) (bool, error) {
	return RDB.HGet(ctx, key, field).Bool()
}

// HGetObject HGet的工具方法，当字段值为bool类型时使用
func HGetObject(key, field string, val interface{}) error {
	reply, err := RDB.HGet(ctx, key, field).Result()
	if err != nil {
		return err
	}
	return conCode(reply, val)
}

// HGetAll 所有的字段和值。
func HGetAll(key string, val interface{}) (map[string]string, error) {
	return RDB.HGetAll(ctx, key).Result()
}

// SisMember 断成员元素是否是集合的成员,如果成员元素是集合的成员，返回 1 。 如果成员元素不是集合的成员，或 key 不存在，返回 0 。
func SisMember(key string, val interface{}) (bool, error) {
	return RDB.SIsMember(ctx, key, val).Result()
}

// Scard 命令返回集合中元素的数量。
func Scard(key string) (int64, error) {
	return RDB.SCard(ctx, key).Result()
}

// SAdd 将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略。
func SAdd(key string, val interface{}, expire int64) error {
	pipe := RDB.TxPipeline()
	pipe.SAdd(ctx, key, val)
	pipe.Expire(ctx, key, time.Duration(expire))
	_, err := pipe.Exec(ctx)
	return err
}

// BLPop 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表
// 直到等待超时或发现可弹出元素为止。
func BLPopString(timeout int64, key ...string) (string, error) {
	values, err := RDB.BLPop(ctx, time.Duration(timeout), key...).Result()
	if err != nil {
		return "", err
	}
	if len(values) != 2 {
		return "", fmt.Errorf("goredis: unexpected number of values, got %d", len(values))
	}
	return values[1], nil
}

// BLPopInt  BLPopInt BLPop的工具方法，元素类型为int时
func BLPopInt(timeout int64, key ...string) (int, error) {
	reply, err := BLPopString(timeout, key...)
	if err != nil {
		return 0, err
	}
	value, _ := strconv.Atoi(reply)
	return value, nil
}

// BLPopInt64 BLPop的工具方法，元素类型为int64时
func BLPopInt64(timeout int64, key ...string) (int64, error) {
	reply, err := BLPopString(timeout, key...)
	if err != nil {
		return 0, err
	}
	value, _ := strconv.ParseInt(reply, 10, 64)
	return value, nil
}

// BLPopBool BLPop的工具方法，元素类型为bool时
func BLPopBool(timeout int64, key ...string) (bool, error) {
	reply, err := BLPopString(timeout, key...)
	if err != nil {
		return false, err
	}
	value, _ := strconv.ParseBool(reply)
	return value, nil
}

// BLPopObject BLPop的工具方法，元素类型为object时
func BLPopObject(timeout int64, val interface{}, key ...string) error {
	reply, err := BLPopString(timeout, key...)
	if err != nil {
		return err
	}
	return conCode(reply, val)
}

// BRPop 命令移出并获取列表的最后一个元素，
// 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
func BRPopString(timeout int64, key ...string) (string, error) {
	values, err := RDB.BRPop(ctx, time.Duration(timeout), key...).Result()
	if err != nil {
		return "", err
	}
	if len(values) != 2 {
		return "", fmt.Errorf("goredis: unexpected number of values, got %d", len(values))
	}
	return values[1], nil
}

// BRPopInt BRPop的工具方法，元素类型为int时
func BRPopInt(timeout int64, key ...string) (int, error) {
	reply, err := BLPopString(timeout, key...)
	if err != nil {
		return 0, err
	}
	value, _ := strconv.Atoi(reply)
	return value, nil
}

// BRPopInt64 BRPop的工具方法，元素类型为int64时
func BRPopInt64(timeout int64, key ...string) (int64, error) {
	reply, err := BLPopString(timeout, key...)
	if err != nil {
		return 0, err
	}
	value, _ := strconv.ParseInt(reply, 10, 64)
	return value, nil
}

// BRPopBool BRPop的工具方法，元素类型为bool时
func BRPopBool(timeout int64, key ...string) (bool, error) {
	reply, err := BLPopString(timeout, key...)
	if err != nil {
		return false, err
	}
	value, _ := strconv.ParseBool(reply)
	return value, nil
}

// BRPopObject BRPop的工具方法，元素类型为object时
func BRPopObject(timeout int64, val interface{}, key ...string) error {
	reply, err := BLPopString(timeout, key...)
	if err != nil {
		return err
	}
	return conCode(reply, val)
}

// LLen 用于返回列表的长度。 如果列表 key 不存在，则 key 被解释为一个空列表，返回 0 。
// 如果 key 不是列表类型，返回一个错误。
func LLen(key string) (int64, error) {
	return RDB.LLen(ctx, key).Result()
}

// LPop 移出并获取列表中第一个元素
func LPopString(key string) (string, error) {
	return RDB.LPop(ctx, key).Result()
}

// LPopInt 移出并获取列表中的第一个元素，元素类型为int
func LPopInt(key string) (int, error) {
	return RDB.LPop(ctx, key).Int()
}

// LPopInt64 移出并获取列表中的第一个元素，元素类型为int64
func LPopInt64(key string) (int64, error) {
	return RDB.LPop(ctx, key).Int64()
}

// LPopBool 移出并获取列表中的第一个元素，元素类型为bool
func LPopBool(key string) (bool, error) {
	return RDB.LPop(ctx, key).Bool()
}

// LPopObject 移出并获取列表中的第一个元素（表头，左边），元素类型为非基本类型的struct
func LPopObject(key string, val interface{}) error {
	reply, err := GetString(key)
	if err != nil {
		return err
	}
	return conCode(reply, val)
}

// LPush  将一个值插入到列表头部
func LPush(key string, val interface{}) error {
	_, err := RDB.LPush(ctx, key, val).Result()
	return err
}

// RPush 将一个值插入到列表尾部
func RPush(key string, val interface{}) error {
	_, err := RDB.RPush(ctx, key, val).Result()
	return err
}

// ZCard 获取有序集合的成员数
func ZCard(key string) (int64, error) {
	return RDB.ZCard(ctx, key).Result()
}

// ZScore 命令返回有序集中，成员的分数值。
func ZScore(key, member string) (float64, error) {
	return RDB.ZScore(ctx, key, member).Result()
}

// conCode 反序列化数据
func conCode(reply string, val interface{}) error {
	return json.Unmarshal([]byte(reply), val)
}

// errorHandle 错误打印
func errorHandle(err error) {
	if err == redis.Nil {
		return
	}
	logrus.Debug("", map[string]interface{}{
		"error": fmt.Sprint(err),
		"stack": string(debug.Stack()),
	})
}
