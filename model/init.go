package model

import (
	"fmt"
	"gin/config"
	"gin/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// Database ...
func Init(config config.DatabaseConfig) error {
	//dsn="root:admin@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
			TablePrefix:   "cmf_",
		},
		//禁用默认事务，防止重复提交/回滚
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("连接数据库异常,err", err)
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	//最大闲置连接数
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	//最大的连接数，默认值为0表示不限制
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	// Auto()
	// DB.AutoMigrate(
	// 	&Pushrecord{},
	// )
	log.Info("数据库连接成功！")
	return nil
}
