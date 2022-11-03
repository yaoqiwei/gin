package config

import (
	"strconv"

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
func Api() ApiConfig {
	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}
	return ApiConfig{
		Port: port,
	}
}

// Database : 数据库配置内容
func Database() DatabaseConfig {
	key := "DatabaseConfig."
	Host := viper.GetString(key + "Host")
	if Host == "" {
		Host = "127.0.0.1"
	}
	Port := viper.GetString(key + "Port")
	if Port == "" {
		Port = "3306"
	}
	User := viper.GetString(key + "User")
	Password := viper.GetString(key + "Password")
	Name := viper.GetString(key + "Name")
	Charset := viper.GetString(key + "CHARSET")
	if Charset == "" {
		Charset = "utf8mb4"
	}
	ShowSQL := viper.GetString("DB_SHOW_SQL")
	MaxIdleConns, err := strconv.Atoi(viper.GetString("DB_MAX_IDLE_CONNS"))
	if err != nil {
		MaxIdleConns = 2
	}
	MaxOpenConns, err := strconv.Atoi(viper.GetString("DB_MAX_OPEN_CONNS"))
	if err != nil {
		MaxOpenConns = 10
	}
	return DatabaseConfig{
		Host:         Host,
		Port:         Port,
		User:         User,
		Password:     Password,
		Name:         Name,
		Charset:      Charset,
		ShowSQL:      ShowSQL == "true",
		MaxIdleConns: MaxIdleConns,
		MaxOpenConns: MaxOpenConns,
	}
}

/*Redis : Redis 配置内容*/
func Redis() RedisConfig {
	key := "REDIS"
	Host := viper.GetString(key + "HOST")
	if Host == "" {
		Host = "127.0.0.1"
	}
	Port := viper.GetString(key + "PORT")
	if Port == "" {
		Port = "6379"
	}
	Password := viper.GetString(key + "PASSWORD")
	DB := viper.GetInt(key + "DB")
	return RedisConfig{
		Host:     Host,
		Port:     Port,
		Password: Password,
		DB:       DB,
	}
}
