package gorm

import (
	"go-boot/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
func GormMysql(c *config.Config) *gorm.DB {
	m := c.Mysql
	if m.Dbname == "" {
		return nil
	}
	config := mysql.Config{
		// 连接数据库DSN
		DSN: m.Dsn(),
		// string 类型字段的默认长度
		DefaultStringSize: 191,
		// 根据版本自动配置
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(config), Config(c))
	if err != nil {
		return nil
	}
	db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
	sqlDB, _ := db.DB()
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	// 最大连接数
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	// 连接最大生存时间
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
	return db
}

// GormMysqlByConfig 根据外部配置实例化*gorm.DB
func GormMysqlByConfig(c *config.Config) *gorm.DB {
	if c.Mysql.Dbname == "" {
		return nil
	}
	config := mysql.Config{
		// 连接数据库DSN
		DSN: c.Mysql.Dsn(),
		// string 类型字段的默认长度
		DefaultStringSize: 191,
		// 根据版本自动配置
		SkipInitializeWithVersion: false,
	}

	if db, err := gorm.Open(mysql.New(config), Config(c)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.Mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.Mysql.MaxOpenConns)
		return db
	}
}
