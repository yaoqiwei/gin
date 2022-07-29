package main

import (
	"gin/config"
	"gin/log"
	"gin/model"
	"gin/redis"
	"gin/routes"
)

func main() {

	logConfig := config.Log()
	DatabaseConfig := config.Database()
	RedisConfig := config.Redis()

	log.Init(logConfig)
	model.Init(DatabaseConfig)
	redis.Init(RedisConfig)

	routes.InitRouter()
}
