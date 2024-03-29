package gorm

import (
	"errors"
	"gin/config"
	"gin/config/structs"
	"gin/util/stringify"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GormGinPool map[string]*gorm.DB
var GinDb *DbWrapper

type DbWrapper struct {
	*gorm.DB
	Split int64
}

// SubTable 分表设置
func (d *DbWrapper) SubTable(tableName string, i interface{}) *gorm.DB {
	if d.Split == 0 {
		return d.DB
	}
	v := stringify.ToInt(i)
	suffix := "_" + stringify.ToString(v%d.Split)
	return d.Table(tableName + suffix)
}

// SubTable 分表设置
func SubTable(d *DbWrapper, tableName string, i interface{}) string {
	if d.Split == 0 {
		return tableName
	}
	v := stringify.ToInt(i)
	suffix := "_" + stringify.ToString(v%d.Split)
	return tableName + suffix
}

// OutTimePrint sql执行时间超出打印
func (d *DbWrapper) OutTimePrint(t int) *gorm.DB {
	return d.Session(&gorm.Session{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Duration(t) * time.Second, // 慢sql阈值
			LogLevel:                  logger.Warn,                    // 日志级别
			IgnoreRecordNotFoundError: true,                           // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                          // 禁用彩色打印
		}),
	})
}

// InitGormPool 配置gorm
func InitGormPool() error {
	//视频数据源
	GormGinPool = map[string]*gorm.DB{}
	err := setDbPoll(GormGinPool, config.MysqlGin)
	if err != nil {
		return err
	}
	GinDb, err = GetGinPool("default")
	if err != nil {
		return err
	}
	return nil
}

// setDbPoll 设置数据库连接池
func setDbPoll(dbPool map[string]*gorm.DB, mysqlConf structs.MysqlConf) error {
	for confName, DbConf := range mysqlConf.List {
		gormDB, err := gorm.Open(mysql.Open(DbConf.DataSourceName), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// TablePrefix:   DbConf.Prefix,
				SingularTable: true, // 使用单数表名
			},
			//禁用默认事务，防止重复提交/回滚
			SkipDefaultTransaction: true,
		})
		if err != nil {
			return err
		}
		sqlDB, err := gormDB.DB()
		if err != nil {
			return err
		}
		//最大闲置连接数
		sqlDB.SetMaxIdleConns(DbConf.MaxIdleConn)
		//最大的连接数，默认值为0表示不限制
		sqlDB.SetMaxOpenConns(DbConf.MaxOpenConn)
		//最大连接超时
		sqlDB.SetConnMaxLifetime(time.Duration(DbConf.MaxConnLifeTime) * time.Second)

		err = sqlDB.Ping()
		if err != nil {
			return err
		}

		dbPool[confName] = gormDB
	}
	return nil
}

// GetGinPool 获取gorm连接池
func GetGinPool(name string) (*DbWrapper, error) {
	if dbPool, ok := GormGinPool[name]; ok {
		dbWrapper := DbWrapper{}
		dbWrapper.DB = dbPool
		dbWrapper.Split = int64(config.MysqlGin.Split)
		return &dbWrapper, nil
	}
	return nil, errors.New("get gormPool error")
}

// BlendOr 混合OR查询条件
func BlendOr(index int, query string) string {
	if index > 0 {
		query = " OR " + query
	}
	return query
}
