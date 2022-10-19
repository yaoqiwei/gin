package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

/*Init : 初始化配置*/
func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Errorf("no such .env file, use system env", err)
	} else {
		logrus.Info("use .env file")
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
