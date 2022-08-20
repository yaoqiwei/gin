package config

const envPerfix string = "DVC_"

/*LogConfig : 日志配置文件结构*/
type LogConfig struct {
	Filename string
	Console  bool
}

/*ApiConfig : API配置文件结构*/
type ApiConfig struct {
	Port string
}

/*DatabaseConfig : 数据库配置文件结构*/
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

/*RedisConfig : REDIS配置文件结构*/
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

/*CaptchaConfig : 验证码校正*/
type CaptchaConfig struct {
	Status             bool
	Url                string
	Login_Aid          string
	Login_AppSecretKey string
	Sms_Aid            string
	Sms_AppSecretKey   string
}

/*SmsConfig : 短信发送配置 */
type SmsConfig struct {
	Status       bool
	Url          string
	TimestampUrl string
	Appid        string
	Appkey       string
	Reg          string
	Login        string
}

/* ObsConfig : Obs配置 */
type ObsConfig struct {
	Access  string
	Secret  string
	Server  string
	Bucket  string
	Domain  string
	TempDir string // 临时文件目录
}

/* AgencyConfig : 易代理配置 */
type AgencyConfig struct {
	AgencyKey string
	Action    string
	Qty       string
	OrderNum  string
	Isp       string
}

/* AdminConfig : 管理员配置 */
type AdminConfig struct {
	Username string
	Password string
}

/* KafkaConfig : Kafka配置 */
type KafkaConfig struct {
	KafkaHost string
}
