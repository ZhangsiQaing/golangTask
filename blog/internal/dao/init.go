package dao

import (
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) error {
	// return db.AutoMigrate(&UserDAO{}, &PostDAO{}, &CommentDAO{})
	return db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
