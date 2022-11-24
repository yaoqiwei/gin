package config

import (
	"gin/config/structs"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Base structs.BaseConfig
var Http structs.HttpConf
var TimeZone *time.Location
var MysqlGin structs.MysqlConf
var RedisConf structs.RedisConfig
var RabbitMQConf structs.RabbitMQConfig
var DebugMode string

// Init : 初始化配置
func init() {
	viper.SetConfigFile("./gin_config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("fatal error config file:", err)
	}
	viper.Unmarshal(&Base)

}

// Api : API配置内容
func HttpConf() {
	Http = Base.HttpConf
	if Http.Addr == "" {
		Http.Addr = "8080"
	}
}

// Database : 数据库配置内容
func Database() {
	MysqlGin = Base.MysqlConf
}

// ModeConf : ModeConf 配置
func ModeConf() {
	DebugMode = Base.DebugMode
	if DebugMode == "" {
		DebugMode = "dev"
	}
}

// Redis : Redis 配置内容
func Redis() {
	RedisConf = Base.RedisConfig

	if RedisConf.Host == "" {
		RedisConf.Host = "127.0.0.1"
	}

	if RedisConf.Port == "" {
		RedisConf.Port = "6379"
	}
}

// RabbitMQ : RabbitMQ 配置
func RabbitMQ() {
	RabbitMQConf = Base.RabbitMQConfig
	if RabbitMQConf.Host == "" {
		RabbitMQConf.Host = "127.0.0.1"
	}
	if RabbitMQConf.Port == "" {
		RabbitMQConf.Port = "15672"
	}
	if RabbitMQConf.UserName == "" {
		RabbitMQConf.UserName = "guest"
	}
	if RabbitMQConf.PassWord == "" {
		RabbitMQConf.PassWord = "guest"
	}
}
