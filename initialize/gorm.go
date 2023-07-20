package initialize

import (
	"go-boot/global"
	orm "go-boot/initialize/gorm"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	dbType := global.CONFIG.System.DbType
	dbMaps := map[string]*gorm.DB{
		"mysql": orm.GormMysql(),
	}
	if db, exists := dbMaps[dbType]; exists {
		return db
	}
	return dbMaps["mysql"]
}
