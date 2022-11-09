package controller

import (
	"fmt"
	"gin/context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

type CornTestService struct {
	Pushtime string `json:"pushtime"`
}

func CornTest(c *gin.Context) {
	p := CornTestService{}

	if err := c.ShouldBind(&p); err != nil {
		context.ValidateError(c)
		return
	}
	pushtime, _ := time.ParseInLocation("2006-01-02 15:04:05", p.Pushtime, time.Local)
	today := time.Now()
	mistiming := pushtime.Sub(today)
	logrus.Println("Starting...")
	a := cron.New()
	fmt.Println(mistiming.String())
	a.AddFunc("0 */1 * * * * ", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})
	a.Start()
	defer a.Stop()
	select {}
}
