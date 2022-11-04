package main

import (
	"gin/config"
	"gin/log"
	"gin/redis"
	"gin/routes"
)

func main() {
	RedisConfig := config.Redis()
	config.Database()

	log.Init()
	redis.Init(RedisConfig)

	routes.HttpServerRun()

}
