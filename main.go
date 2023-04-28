package main

import (
	"gin/common/lib"
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
	config.ModeConf()

	lib.Init()
	// rabbitmq.InitRabbitMQ()
	routes.HttpServerRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
