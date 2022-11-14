package controller

import (
	"fmt"
	"gin/common/lib"
	"gin/common/lib/rabbitmq"
	"time"

	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func TestRegister(router lib.RegisterRouterGroup, needLoginRouter lib.RegisterRouterGroup) {
	c := TestController{}
	needLoginRouter.POST("/test", c.CornTest)
	needLoginRouter.POST("/send", c.SendTest)
}

type CornTestService struct {
	Pushtime string `json:"pushtime"`
}

func (*TestController) CornTest(c *gin.Context) {
	bufChan := make(chan int, 5)

	go func() {
		time.Sleep(time.Second)
		for {
			<-bufChan
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case bufChan <- 1:
			fmt.Println("add success")
			time.Sleep(time.Second)
		default:
			fmt.Println("资源已满，请稍后再试")
			time.Sleep(time.Second)
		}
	}
}

func (*TestController) SendTest(c *gin.Context) {
	fmt.Println("AAA", rabbitmq.RabbitMqConn)
	for i := 0; i < 100; i++ {
		rabbitmq.RabbitMqConn.SendMessage([]byte{12}, "aaa")
	}
}
