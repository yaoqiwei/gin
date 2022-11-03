package config

// ApiConfig : API配置文件结构
type ApiConfig struct {
	Port string
}

// DatabaseConfig : 数据库配置文件结构
type DatabaseConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	Charset      string
	ShowSQL      bool
	MaxIdleConns int
	MaxOpenConns int
}

// RedisConfig : REDIS配置文件结构
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}
