package log

import (
	"gin/config"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

/*Init : 初始化日志*/
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
}
