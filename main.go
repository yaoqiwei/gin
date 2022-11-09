package main

import (
	"gin/common/lib"
	"gin/config"
	"gin/routes"
)

func main() {
	config.HttpConf()
	config.Database()
	config.Redis()

	lib.Init()
	routes.HttpServerRun()
}
