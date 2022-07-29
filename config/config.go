package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

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

/*Init : 初始化配置*/
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no such .env file, use system env")
	} else {
		log.Println("use .env file")
	}
}

/*Log : 日志配置内容*/
func Log() LogConfig {
	logConsole := Getenv("LOG_CONSOLE")
	logFilename := Getenv("LOG_FILENAME")
	if logFilename == "" {
		logFilename = "run"
	}
	return LogConfig{
		Filename: logFilename,
		Console:  logConsole == "true",
	}
}

/*Api : API配置内容*/
func Api() ApiConfig {
	port := Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	return ApiConfig{
		Port: port,
	}
}

/*Database : 数据库配置内容*/
func Database() DatabaseConfig {
	Host := Getenv("DB_HOST")
	if Host == "" {
		Host = "127.0.0.1"
	}
	Port := Getenv("DB_PORT")
	if Port == "" {
		Port = "3306"
	}
	User := Getenv("DB_USER")
	Password := Getenv("DB_PASSWORD")
	Name := Getenv("DB_NAME")
	Charset := Getenv("DB_CHARSET")
	if Charset == "" {
		Charset = "utf8mb4"
	}
	ShowSQL := Getenv("DB_SHOW_SQL")
	MaxIdleConns, err := strconv.Atoi(Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		MaxIdleConns = 2
	}
	MaxOpenConns, err := strconv.Atoi(Getenv("DB_MAX_OPEN_CONNS"))
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
	Host := Getenv("REDIS_HOST")
	if Host == "" {
		Host = "127.0.0.1"
	}
	Port := Getenv("REDIS_PORT")
	if Port == "" {
		Port = "6379"
	}
	Password := Getenv("REDIS_PASSWORD")
	DB, err := strconv.Atoi(Getenv("REDIS_DB"))
	if err != nil {
		DB = 0
	}
	return RedisConfig{
		Host:     Host,
		Port:     Port,
		Password: Password,
		DB:       DB,
	}
}

/*Captcha : Captcha 配置内容*/
func Captcha() CaptchaConfig {
	Status := true
	Url := Getenv("CAPTCHA_URL")
	Login_Aid := Getenv("CAPTCHA_LOGIN_AID")
	Login_AppSecretKey := Getenv("CAPTCHA_LOGIN_APPSECRETKEY")
	Sms_Aid := Getenv("CAPTCHA_SMS_AID")
	Sms_AppSecretKey := Getenv("CAPTCHA_SMS_APPSECRETKEY")
	return CaptchaConfig{
		Status,
		Url,
		Login_Aid,
		Login_AppSecretKey,
		Sms_Aid,
		Sms_AppSecretKey,
	}
}

/*Sms : Sms配置内容*/
func Sms() SmsConfig {
	Status := true
	Url := Getenv("SMS_URL")
	TimestampUrl := Getenv("SMS_TIMESTAMPURL")
	Appid := Getenv("SMS_APPID")
	Appkey := Getenv("SMS_APPKEY")
	Reg := Getenv("SMS_REG")
	Login := Getenv("SMS_LOGIN")
	return SmsConfig{
		Status,
		Url,
		TimestampUrl,
		Appid,
		Appkey,
		Reg,
		Login,
	}
}

/* Admin : Admin配置 */
func Admin() AdminConfig {
	Username := Getenv("ADMIN_USER")
	if Username == "" {
		Username = "admin"
	}
	Password := Getenv("ADMIN_PASSWORD")
	if Password == "" {
		Password = "888888"
	}
	return AdminConfig{
		Username: Username,
		Password: Password,
	}
}

/* Excel : Excel配置 */
func Excel() string {
	ExcelDir := Getenv("EXCEL_DIR")
	if ExcelDir == "" {
		ExcelDir = "excel"
	}
	return ExcelDir
}

/* Obs : Obs配置 */
func Obs() ObsConfig {
	accessKeyId := Getenv("UPLOAD_OBS_ACCESS_KEY_ID")
	secretAccessKey := Getenv("UPLOAD_OBS_SECRET_ACCESS_KEY")
	server := Getenv("UPLOAD_OBS_SERVER")
	bucket := Getenv("UPLOAD_OBS_BUCKET")
	domain := Getenv("UPLOAD_OBS_DOMAIN")
	tempDir := Getenv("UPLOAD_TEMP_DIR")
	if tempDir == "" {
		tempDir = "upload"
	}
	return ObsConfig{
		Access:  accessKeyId,
		Secret:  secretAccessKey,
		Server:  server,
		Bucket:  bucket,
		Domain:  domain,
		TempDir: tempDir,
	}
}

/*Getenv : 获取环境变量*/
func Getenv(name string) string {
	return os.Getenv(envPerfix + name)
}
