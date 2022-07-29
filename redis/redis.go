package redis

import (
	"context"
	"gin/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RDB *redis.Client
var DefaultExpiration = time.Minute * 1440 * 7

/*Init : 初始化REDIS*/
func Init(config config.RedisConfig) {
	addr := config.Host + ":" + config.Port
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DB,
	})
}
