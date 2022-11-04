package structs

import "time"

// ApiConfig : API配置文件结构
type ApiConfig struct {
	Port string
}

// http配置
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

// MysqlConf 数据库配置文件
type MysqlConf struct {
	DriverName      string `yaml:"DriverName"`
	DataSourceName  string `yaml:"DataSourceName"`
	MaxOpenConn     int    `yaml:"MaxOpenConn"`
	MaxIdleConn     int    `yaml:"MaxIdleConn"`
	MaxConnLifeTime int    `yaml:"MaxConnLifeTime"`
	Prefix          string `yaml:"Prefix"`
}

// MysqlMapConfig
type MysqlMapConfig struct {
	List  map[string]*MysqlConf `yaml:"List"`
	Split int                   `yaml:"Split"`
}

// RedisConfig : REDIS配置文件结构
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}
