package main

import (
	"gin/config"
	"gin/log"
	"gin/model"
	"gin/redis"
	"gin/routes"
)

func main() {
	DatabaseConfig := config.Database()
	RedisConfig := config.Redis()

	log.Init()
	model.Init(DatabaseConfig)
	redis.Init(RedisConfig)

	routes.HttpServerRun()

}
