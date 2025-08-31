package sql

import (
	"fmt"

	"gorm.io/gorm"
)

func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	if err := db.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_number", gorm.Expr("post_number + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// AfterDelete 钩子
func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
	var count int64
	fmt.Println(c)
	if err := db.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		fmt.Println(c.PostID)
		// 如果评论数量为 0，更新文章状态
		if err := db.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("status", "无评论").Error; err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(err)
	}
	return nil
}
