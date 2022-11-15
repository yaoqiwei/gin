package lib

import (
	"gin/common/lib/gorm"
	"gin/common/lib/redis"
	"gin/config"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func Init() {
	time.Local = config.TimeZone
	color.NoColor = false

	logger := logrus.StandardLogger()
	logger.Out = os.Stdout

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006/01/02 - 15:04:05.000",
	})

	logrus.SetLevel(logrus.TraceLevel)

	redis.Init()

	if err := gorm.InitGormPool(); err != nil {
		logrus.Errorf("InitGromPool:" + err.Error())
	}
	logrus.Infof("success loading resources.")
	logrus.Infof("------------------------------------------------------------------------")
}

// Destroy 公共销毁函数
func Destroy() {
	logrus.Infof("------------------------------------------------------------------------")
	logrus.Infof("start Destroying resources.")
	gorm.CloseDB()
	logrus.Infof("success Destroying resources.")
}
