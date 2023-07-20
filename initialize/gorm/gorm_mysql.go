package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"go-boot/config"
	"go-boot/global"
	"go-boot/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.CONFIG.Mysql
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

	if db, err := gorm.Open(mysql.New(config), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		// 设置最大空闲连接数
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		// 设置最大连接数
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMysqlByConfig 根据外部配置实例化*gorm.DB
func GormMysqlByConfig(m config.Mysql) *gorm.DB {
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
	if db, err := gorm.Open(mysql.New(config), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
