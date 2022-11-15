package main

import (
	"gin/common/lib"
	"gin/common/lib/rabbitmq"
	"gin/config"
	"gin/routes"
	"os"
	"os/signal"
)

func main() {
	config.HttpConf()
	config.Database()
	config.Redis()
	config.RabbitMQ()

	lib.Init()
	defer lib.Destroy()
	rabbitmq.InitRabbitMQ()
	routes.HttpServerRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
