package gorm

import (
	"go-boot/config"
	"gorm.io/gorm"
)

// NewGorm 初始化Gorm
func NewGorm(config *config.Config) *gorm.DB {
	// 获取DB类型
	dbType := config.System.DbType
	dbMaps := map[string]*gorm.DB{
		"mysql": GormMysql(config),
	}
	if db, exists := dbMaps[dbType]; exists {
		return db
	}
	return dbMaps["mysql"]
}
