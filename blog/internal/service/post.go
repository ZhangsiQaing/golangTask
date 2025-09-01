package service

import (
	"blog/internal/domain"
	"blog/internal/repository"
	"context"
	"errors"
)

var (
	ErrPostNotFound   = errors.New("文章不存在")
	ErrPostPermission = errors.New("权限不足，只有文章作者才能操作")
	ErrInvalidPostID  = errors.New("无效的文章ID")
	ErrInvalidUserID  = errors.New("无效的用户ID")
	ErrEmptyTitle     = errors.New("文章标题不能为空")
	ErrEmptyContent   = errors.New("文章内容不能为空")
	ErrTitleTooLong   = errors.New("文章标题长度不能超过255个字符")
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

// CreatePost 创建文章
func (s *PostService) CreatePost(ctx context.Context, req domain.CreatePostRequest, authorId uint) error {
	// 参数验证
	if req.Title == "" {
		return ErrEmptyTitle
	}

	if req.Content == "" {
		return ErrEmptyContent
	}

	if len(req.Title) > 255 {
		return ErrTitleTooLong
	}

	return s.postRepo.Create(ctx, req, authorId)
}

// GetPost 获取单个文章详情
func (s *PostService) GetPost(ctx context.Context, id uint) (domain.PostDetailResponse, error) {
	if id == 0 {
		return domain.PostDetailResponse{}, ErrInvalidPostID
	}

	// 检查文章是否存在
	exists, err := s.postRepo.Exists(ctx, id)
	if err != nil {
		return domain.PostDetailResponse{}, err
	}

	if !exists {
		return domain.PostDetailResponse{}, ErrPostNotFound
	}

	return s.postRepo.GetById(ctx, id)
}

// GetAllPosts 获取所有文章列表
func (s *PostService) GetAllPosts(ctx context.Context) ([]domain.PostListResponse, error) {
	return s.postRepo.GetAll(ctx)
}

// GetPostsByUserId 根据用户ID获取文章列表
func (s *PostService) GetPostsByUserId(ctx context.Context, userId uint) ([]domain.PostListResponse, error) {
	if userId == 0 {
		return nil, ErrInvalidUserID
	}

	return s.postRepo.GetByUserId(ctx, userId)
}

// UpdatePost 更新文章
func (s *PostService) UpdatePost(ctx context.Context, id uint, authorId uint, req domain.UpdatePostRequest) error {
	if id == 0 {
		return ErrInvalidPostID
	}

	if authorId == 0 {
		return ErrInvalidUserID
	}

	// 参数验证
	if req.Title == "" {
		return ErrEmptyTitle
	}

	if req.Content == "" {
		return ErrEmptyContent
	}

	if len(req.Title) > 255 {
		return ErrTitleTooLong
	}

	// 检查文章是否存在
	exists, err := s.postRepo.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return ErrPostNotFound
	}

	// 检查是否为文章作者
	isAuthor, err := s.postRepo.IsAuthor(ctx, id, authorId)
	if err != nil {
		return ErrPostPermission
	}

	if !isAuthor {
		return ErrPostPermission
	}

	return s.postRepo.Update(ctx, id, authorId, req)
}

// DeletePost 删除文章
func (s *PostService) DeletePost(ctx context.Context, id uint, authorId uint) error {
	if id == 0 {
		return ErrInvalidPostID
	}

	if authorId == 0 {
		return ErrInvalidUserID
	}

	// 检查文章是否存在
	exists, err := s.postRepo.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return ErrPostNotFound
	}

	// 检查是否为文章作者
	isAuthor, err := s.postRepo.IsAuthor(ctx, id, authorId)
	if err != nil {
		return err
	}

	if !isAuthor {
		return ErrPostPermission
	}

	return s.postRepo.Delete(ctx, id, authorId)
}
