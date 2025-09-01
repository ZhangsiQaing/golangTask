package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null;index"`
	PostID  uint   `gorm:"not null;index"`
}

type CommentDAO struct {
	db *gorm.DB
}

func NewCommentDAO(db *gorm.DB) *CommentDAO {
	return &CommentDAO{db: db}
}

// Create 创建评论
func (dao *CommentDAO) Create(ctx context.Context, comment Comment) error {
	now := time.Now()
	comment.CreatedAt = now
	comment.UpdatedAt = now
	return dao.db.WithContext(ctx).Create(&comment).Error
}

// FindByPostId 根据文章ID查找评论
func (dao *CommentDAO) FindByPostId(ctx context.Context, postId uint) ([]Comment, error) {
	var comments []Comment
	err := dao.db.WithContext(ctx).Where("post_id = ?", postId).Order("created_at ASC").Find(&comments).Error
	return comments, err
}

// IsAuthor 检查用户是否为评论作者
func (dao *CommentDAO) IsAuthor(ctx context.Context, commentId uint, userId uint) (bool, error) {
	var comment Comment
	err := dao.db.WithContext(ctx).Select("user_id").Where("id = ?", commentId).First(&comment).Error
	if err != nil {
		return false, err
	}
	return comment.UserID == userId, nil
}
