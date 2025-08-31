package sql

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// User 用户
type User struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"size:100;not null"`
	Email      string    `gorm:"size:100;uniqueIndex;not null"`
	PostNumber uint      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`

	PostList []Post
}

// Post 文章
type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"size:200;not null;default:''"` // 默认空字符串
	Content   string    `gorm:"size:200;not null;default:''"` // 默认空字符串
	UserID    uint      `gorm:"not null;default:0"`           // 默认0
	Status    string    `gorm:"size:20;not null;default:''"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`

	CommentList []Comment
}

// Comment 评论
type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Content   string    `gorm:"size:200;not null;default:''"`         // 默认空字符串
	PostID    uint      `gorm:"not null;default:0"`                   // 默认0
	Status    string    `gorm:"type:varchar(20);not null;default:''"` // 默认空字符串
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}

// 查询某个用户发布的所有文章以及对应的评论信息
func QueryAllInfomationByUid(db *gorm.DB, uid uint) (*User, error) {
	var user User
	// 查询用户
	if err := db.Model(&User{}).First(&user, uid).Error; err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 查询该用户的所有文章
	var posts []Post
	if err := db.Where("user_id = ?", uid).Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("查询文章失败: %w", err)
	}

	if len(posts) == 0 {
		user.PostList = []Post{}
		return &user, nil
	}

	// 收集所有文章 ID
	var postIDs []uint
	for _, p := range posts {
		postIDs = append(postIDs, p.ID)
	}

	// 一次性查所有评论
	var comments []Comment
	if err := db.Where("post_id IN ?", postIDs).Find(&comments).Error; err != nil {
		return nil, fmt.Errorf("查询评论失败: %w", err)
	}

	// 把评论按 post_id 分组
	commentMap := make(map[uint][]Comment)
	for _, c := range comments {
		commentMap[c.PostID] = append(commentMap[c.PostID], c)
	}

	// 给每个文章绑定评论
	for i := range posts {
		posts[i].CommentList = commentMap[posts[i].ID]
	}

	user.PostList = posts

	return &user, nil
}

// 查询评论数量最多的文章信息
func QueryMostCommentedPost(db *gorm.DB) (*Post, int64, error) {
	var post Post
	var commentCount int64

	// 用 JOIN + GROUP BY + ORDER
	err := db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&post).Error
	if err != nil {
		return nil, 0, fmt.Errorf("查询失败: %w", err)
	}

	// 再单独查出评论数量
	if err := db.Model(&Comment{}).
		Where("post_id = ?", post.ID).
		Count(&commentCount).Error; err != nil {
		return nil, 0, fmt.Errorf("查询评论数失败: %w", err)
	}
	return &post, commentCount, nil
}
