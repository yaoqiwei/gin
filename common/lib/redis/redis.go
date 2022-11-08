package redis

import (
	"context"
	"gin/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()
var DefaultExpiration = time.Minute * 1440 * 7

/*Init : 初始化REDIS*/
func Init() {
	addr := config.RedisConf.Host + ":" + config.RedisConf.Port
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.RedisConf.Password,
		DB:       config.RedisConf.DB,
	})
}
