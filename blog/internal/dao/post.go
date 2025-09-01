package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null;index"`
	User    User   `gorm:"foreignKey:UserID"`
}

func (Post) TableName() string {
	return "post"
}

type PostDAO struct {
	db *gorm.DB
}

func NewPostDAO(db *gorm.DB) *PostDAO {
	return &PostDAO{db: db}
}

// Create 创建文章
func (dao *PostDAO) Create(ctx context.Context, post Post) error {
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now
	return dao.db.WithContext(ctx).Create(&post).Error
}

// FindById 根据ID查找文章
func (dao *PostDAO) FindById(ctx context.Context, id uint) (Post, error) {
	var post Post
	err := dao.db.WithContext(ctx).Preload("User").Where("id = ?", id).First(&post).Error
	return post, err
}

// FindAll 获取所有文章
func (dao *PostDAO) FindAll(ctx context.Context) ([]Post, error) {
	var posts []Post
	err := dao.db.WithContext(ctx).Preload("User").Order("created_at DESC").Find(&posts).Error
	return posts, err
}

// FindByUserId 根据用户ID查找文章
func (dao *PostDAO) FindByUserId(ctx context.Context, userId uint) ([]Post, error) {
	var posts []Post
	err := dao.db.WithContext(ctx).Preload("User").Where("user_id = ?", userId).Find(&posts).Error
	return posts, err
}

// Update 更新文章
func (dao *PostDAO) Update(ctx context.Context, id uint, userId uint, title, content string) error {
	now := time.Now()
	return dao.db.WithContext(ctx).Model(&Post{}).
		Where("id = ? AND user_id = ?", id, userId).
		Updates(map[string]interface{}{
			"title":      title,
			"content":    content,
			"updated_at": now,
		}).Error
}

// Delete 删除文章
func (dao *PostDAO) Delete(ctx context.Context, id uint, userId uint) error {
	return dao.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userId).Delete(&Post{}).Error
}

// Exists 检查文章是否存在
func (dao *PostDAO) Exists(ctx context.Context, id uint) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&Post{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

// IsAuthor 检查用户是否为文章作者
func (dao *PostDAO) IsAuthor(ctx context.Context, postId uint, userId uint) (bool, error) {
	var post Post
	err := dao.db.WithContext(ctx).Select("user_id").Where("id = ?", postId).First(&post).Error
	if err != nil {
		return false, err
	}
	return post.UserID == userId, nil
}
