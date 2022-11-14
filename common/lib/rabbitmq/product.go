package rabbitmq

import (
	"encoding/json"
	"gin/config/structs"
	"time"

	"github.com/panjf2000/ants"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
)

func dial(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// ProducerReConnect 生产者重连
func (c *Connection) ProducerReConnect() {
closeTag:
	for {
		c.ConnNotifyClose = c.Conn.NotifyClose(make(chan *amqp.Error))
		c.ChNotifyClose = c.Ch.NotifyClose(make(chan *amqp.Error))
		select {
		case connErr := <-c.ConnNotifyClose:
			if connErr != nil {
				logrus.Errorf("rabbitMQ连接异常:%s", connErr.Error())
			}
			// 判断连接是否关闭
			if !c.Conn.IsClosed() {
				if err := c.Conn.Close(); err != nil {
					logrus.Errorf("rabbit连接关闭异常:%s", err.Error())
				}
			}
			//重新连接
			if conn, err := dial(c.URL); err != nil {
				logrus.Errorf("rabbit重连失败:%s", err.Error())
				_, isConnChannelOpen := <-c.ConnNotifyClose
				if isConnChannelOpen {
					close(c.ConnNotifyClose)
				}
				//connection关闭时会自动关闭channel
				ants.Submit(func() { c.InitRabbitMQProducer(false, c.Rabbit) })
				//结束子进程
				break closeTag
			} else { //连接成功
				c.Ch, _ = conn.Channel()
				c.Conn = conn
				logrus.Info("rabbitMQ重连成功")
			}
			// IMPORTANT: 必须清空 Notify，否则死连接不会释放
			for err := range c.ConnNotifyClose {
				println(err)
			}
		case chErr := <-c.ChNotifyClose:
			if chErr != nil {
				logrus.Errorf("rabbitMQ通道连接关闭:%s", chErr.Error())
			}
			// 重新打开一个并发服务器通道来处理消息
			if !c.Conn.IsClosed() {
				ch, err := c.Conn.Channel()
				if err != nil {
					logrus.Errorf("rabbitMQ channel重连失败:%s", err.Error())
					c.ChNotifyClose <- chErr
				} else {
					logrus.Info("rabbitMQ通道重新创建成功")
					c.Ch = ch
				}
			} else {
				_, isConnChannelOpen := <-c.ConnNotifyClose
				if isConnChannelOpen {
					close(c.ConnNotifyClose)
				}
				ants.Submit(func() { c.InitRabbitMQProducer(false, c.Rabbit) })
				break closeTag
			}
			for err := range c.ChNotifyClose {
				println(err)
			}
		case <-c.CloseProcess:
			break closeTag
		}
	}
	logrus.Info("结束旧生产者进程")
}

// InitRabbitMQProducer 初始化生产者
func (c *Connection) InitRabbitMQProducer(isClose bool, rabbitMQConfig structs.RabbitMQConfig) {
	if isClose {
		c.CloseProcess <- true
	}
	c.Rabbit = rabbitMQConfig
	url := "amqp://" + c.Rabbit.UserName + ":" + c.Rabbit.PassWord + "@" + c.Rabbit.Host + ":" + cast.ToString(c.Rabbit.Port) + "/"
	conn, err := dial(url)
	if err != nil {
		logrus.Errorf("rabbitMQ连接异常:%s", err.Error())
		logrus.Info("休息5S,开始重连rabbitMQ生产者")
		time.Sleep(5 * time.Second)
		ants.Submit(func() { c.InitRabbitMQProducer(false, c.Rabbit) })
		return
	}
	defer conn.Close()
	logrus.Info("rabbitMQ生产者连接成功")
	// 打开一个并发服务器通道来处理消息
	ch, err := conn.Channel()
	if err != nil {
		logrus.Errorf("rabbitMQ打开通道异常:%s", err.Error())
		return
	}
	defer ch.Close()
	c.Conn = conn
	c.URL = url
	c.Ch = ch
	c.CloseProcess = make(chan bool, 1)
	c.ProducerReConnect()
	logrus.Info("结束rabbitMQ旧生产者")
}

func (c *Connection) SendMessage(body []byte, queueName string) {
	if c.RabbitProducerMap == nil {
		logrus.Error("未初始化生产者信息")
		return
	}
	if queueName == "" {
		logrus.Error("队列名称不能为空")
		return
	}
	exchangeName := c.RabbitProducerMap[queueName]
	if exchangeName == "" {
		logrus.Error("交换机名称不能为空")

		return
	}
	m := DLXMessage{
		QueueName:   queueName,
		Content:     string(body),
		NotifyCount: 1,
	}
	body, _ = json.Marshal(m)
	// 发布
	err := c.Ch.Publish(
		exchangeName, // exchange 默认模式，exchange为空
		queueName,    // routing key 默认模式路由到同名队列，即是task_queue
		false,        // mandatory
		false,
		amqp.Publishing{
			// 持久性的发布，因为队列被声明为持久的，发布消息必须加上这个（可能不用），但消息还是可能会丢，如消息到缓存但MQ挂了来不及持久化。
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		logrus.Error("rabbitMQ 发送消息失败:" + err.Error())
		return
	}
}
