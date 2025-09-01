package service

import (
	"blog/internal/domain"
	"blog/internal/repository"
	"context"
	"errors"
)

var (
	ErrCommentNotFound   = errors.New("评论不存在")
	ErrCommentPermission = errors.New("权限不足，只有评论作者才能操作")
	ErrInvalidCommentID  = errors.New("无效的评论ID")
)

type CommentService struct {
	commentRepo *repository.CommentRepository
	postRepo    *repository.PostRepository
}

func NewCommentService(commentRepo *repository.CommentRepository, postRepo *repository.PostRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

// CreateComment 创建评论
func (s *CommentService) CreateComment(ctx context.Context, req domain.CreateCommentRequest, authorId uint) error {
	// 参数验证
	if req.Content == "" {
		return ErrEmptyContent
	}

	if req.PostID == 0 {
		return ErrInvalidPostID
	}

	// 检查文章是否存在
	exists, err := s.postRepo.Exists(ctx, req.PostID)
	if err != nil {
		return err
	}

	if !exists {
		return ErrPostNotFound
	}

	return s.commentRepo.Create(ctx, req, authorId)
}

// GetCommentsByPostId 根据文章ID获取评论列表
func (s *CommentService) GetCommentsByPostId(ctx context.Context, postId uint) ([]domain.CommentListResponse, error) {
	if postId == 0 {
		return nil, ErrInvalidPostID
	}

	// 检查文章是否存在
	exists, err := s.postRepo.Exists(ctx, postId)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrPostNotFound
	}

	return s.commentRepo.GetByPostId(ctx, postId)
}
