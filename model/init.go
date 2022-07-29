package model

import (
	"fmt"
	"gin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database ...
func Init(config config.DatabaseConfig) error {
	//dsn="root:admin@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", config.User, config.Password, config.Host, config.Port, config.Name, config.Charset, config.ShowSQL, "Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库异常,err", err)
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	//最大闲置连接数
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	//最大的连接数，默认值为0表示不限制
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	fmt.Println("连接数据库成功！")
	return nil
}
