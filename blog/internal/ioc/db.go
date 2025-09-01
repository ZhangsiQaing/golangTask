package ioc

import (
	"blog/config"
	"blog/internal/dao"
	"blog/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(l logger.LoggerV1) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		l.Error("初始化数据库表失败", logger.Field{Key: "error", Val: err})
		panic(err)
	}
	return db
}
