package structs

import "time"

type BaseConfig struct {
	HttpConf       HttpConf       `yaml:"HttpConf"`
	TimeZone       TimeZone       `yaml:"TimeZone"`
	MysqlConf      MysqlConf      `yaml:"MysqlConf"`
	RedisConfig    RedisConfig    `yaml:"RedisConfig"`
	RabbitMQConfig RabbitMQConfig `yaml:"RabbitMQConfig"`
}

// TimeZone 时间规格
type TimeZone struct {
	Zone string `yaml:"Zone"`
	Time int    `yaml:"Time"`
}

// HttpConf http配置
type HttpConf struct {
	TimeZone       *time.Location
	Addr           string   `yaml:"Addr"`
	AesOpen        bool     `yaml:"AesOpen"`
	ApiAesKey      string   `yaml:"ApiAesKey"`
	HeaderCheck    bool     `yaml:"HeaderCheck"`
	UploadDomain   string   `yaml:"UploadDomain"`
	UploadExec     string   `yaml:"UploadExec"`
	UploadAuth     string   `yaml:"UploadAuth"`
	TrustedProxies []string `yaml:"TrustedProxies"`
}

// List 数据库配置文件
type List struct {
	DriverName      string `yaml:"DriverName"`
	DataSourceName  string `yaml:"DataSourceName"`
	MaxOpenConn     int    `yaml:"MaxOpenConn"`
	MaxIdleConn     int    `yaml:"MaxIdleConn"`
	MaxConnLifeTime int    `yaml:"MaxConnLifeTime"`
	Prefix          string `yaml:"Prefix"`
}

// MysqlConf
type MysqlConf struct {
	List  map[string]*List `yaml:"List"`
	Split int              `yaml:"Split"`
}

// RedisConfig : REDIS配置文件结构
type RedisConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}

// RabbitMQConfig : RabbitMQ配置文件结构
type RabbitMQConfig struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	UserName string `yaml:"UserName"`
	PassWord string `yaml:"PassWord"`
}
