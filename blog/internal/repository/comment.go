package repository

import (
	"blog/internal/dao"
	"blog/internal/domain"
	"context"
)

type CommentRepository struct {
	commentDAO *dao.CommentDAO
	postDAO    *dao.PostDAO
	userDAO    *dao.UserDAO
}

func NewCommentRepository(commentDAO *dao.CommentDAO, postDAO *dao.PostDAO, userDAO *dao.UserDAO) *CommentRepository {
	return &CommentRepository{
		commentDAO: commentDAO,
		postDAO:    postDAO,
		userDAO:    userDAO,
	}
}

// Create 创建评论
func (r *CommentRepository) Create(ctx context.Context, req domain.CreateCommentRequest, authorId uint) error {
	comment := dao.Comment{
		Content: req.Content,
		UserID:  authorId,
		PostID:  req.PostID,
	}
	return r.commentDAO.Create(ctx, comment)
}

// GetByPostId 根据文章ID获取评论列表
func (r *CommentRepository) GetByPostId(ctx context.Context, postId uint) ([]domain.CommentListResponse, error) {
	comments, err := r.commentDAO.FindByPostId(ctx, postId)
	if err != nil {
		return nil, err
	}

	var responses []domain.CommentListResponse
	for _, comment := range comments {
		responses = append(responses, domain.CommentListResponse{
			ID:      comment.ID,
			Content: comment.Content,
			UserID:  comment.UserID,
			PostID:  comment.PostID,
			// Author:    comment.User.Username,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	return responses, nil
}

// IsAuthor 检查用户是否为评论作者
func (r *CommentRepository) IsAuthor(ctx context.Context, commentId uint, userId uint) (bool, error) {
	return r.commentDAO.IsAuthor(ctx, commentId, userId)
}
