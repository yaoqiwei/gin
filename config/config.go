package config

import (
	"fmt"
	"gin/config/structs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/*Init : 初始化配置*/
func init() {
	viper.SetConfigFile("./gin_config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("fatal error config file:", err)
	}
}

/*Api : API配置内容*/
func Api() structs.ApiConfig {
	port := viper.GetString("PORT")
	fmt.Println("port", port)
	if port == "" {
		port = "8080"
	}
	return structs.ApiConfig{
		Port: port,
	}
}

// Database : 数据库配置内容
func Database() (data structs.MysqlConf) {
	viper.Unmarshal(&data)
	fmt.Println("data:", data)
	return
}

/*Redis : Redis 配置内容*/
func Redis() structs.RedisConfig {
	key := "RedisConfig."
	Host := viper.GetString(key + "HOST")
	fmt.Println("host", Host)
	if Host == "" {
		Host = "127.0.0.1"
	}
	Port := viper.GetString(key + "PORT")
	if Port == "" {
		Port = "6379"
	}
	Password := viper.GetString(key + "PASSWORD")
	DB := viper.GetInt(key + "DB")
	return structs.RedisConfig{
		Host:     Host,
		Port:     Port,
		Password: Password,
		DB:       DB,
	}
}
