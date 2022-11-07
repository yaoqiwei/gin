package routes

import (
	"gin/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HttpServerRun() {
	// 设置模式 1.debug 2.release 3.test
	gin.SetMode(gin.ReleaseMode)
	r := InitRouter()

	addr := config.Http.Addr
	go func() {
		logrus.Infof("HttpServerRun:%s", addr)
		if err := r.Run(addr); err != nil {
			logrus.Errorf("HttpServerRun:%s err:%v", addr, err)
		}
	}()
}
