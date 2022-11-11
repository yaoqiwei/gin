package rabbitmq

import (
	"gin/config"
	"gin/config/structs"

	"github.com/streadway/amqp"
)

var RabbitMqConn Connection

type RabbitConsumerInfo struct {
	ExchangeName string // 交换机
	QueueName    string // 队列
	ExchangeType string // 模式
}

// Rabbit连接
type Connection struct {
	Conn               *amqp.Connection // 连接
	Ch                 *amqp.Channel    // 通道
	ConnNotifyClose    chan *amqp.Error // 连接异常结束
	ChNotifyClose      chan *amqp.Error // 通道异常接收
	URL                string
	Rabbit             structs.RabbitMQConfig
	CloseProcess       chan bool                  // 用于关闭进程
	RabbitConsumerList []RabbitConsumerInfo       // 消费者信息
	RabbitProducerMap  map[string]string          // 生产者信息
	ConsumeHandle      func(<-chan amqp.Delivery) // 自定义消费者处理函数
}

const (
	CancelOrderDelayQueue = "soesoft.cancel.order.delay.queue"
	DelayExchange         = "soesoft.delay.exchange"
)

type DLXMessage struct {
	QueueName   string `json:"queueName"`
	Content     string `json:"content"`
	NotifyCount int    `json:"notifyCount"`
}

func InitRabbitMQ() {

	// 启动消费者
	go func() {
		RabbitMqConn.InitRabbitMQProducer(false, config.RabbitMQConf)
		RabbitMqConn.InitRabbitMQConsumer(false, config.RabbitMQConf)
	}()
}
