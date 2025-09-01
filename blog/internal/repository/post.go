package repository

import (
	"blog/internal/dao"
	"blog/internal/domain"
	"context"
)

type PostRepository struct {
	postDAO *dao.PostDAO
	userDAO *dao.UserDAO
}

func NewPostRepository(postDAO *dao.PostDAO, userDAO *dao.UserDAO) *PostRepository {
	return &PostRepository{
		postDAO: postDAO,
		userDAO: userDAO,
	}
}

// Create 创建文章
func (r *PostRepository) Create(ctx context.Context, req domain.CreatePostRequest, authorId uint) error {
	post := dao.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  authorId,
	}
	return r.postDAO.Create(ctx, post)
}

// GetById 根据ID获取文章详情
func (r *PostRepository) GetById(ctx context.Context, id uint) (domain.PostDetailResponse, error) {
	post, err := r.postDAO.FindById(ctx, id)
	if err != nil {
		return domain.PostDetailResponse{}, err
	}

	return domain.PostDetailResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		Author:    post.User.Username,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

// GetAll 获取所有文章列表
func (r *PostRepository) GetAll(ctx context.Context) ([]domain.PostListResponse, error) {
	posts, err := r.postDAO.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []domain.PostListResponse
	for _, post := range posts {
		responses = append(responses, domain.PostListResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserID:    post.UserID,
			Author:    post.User.Username,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return responses, nil
}

// GetByUserId 根据用户ID获取文章列表
func (r *PostRepository) GetByUserId(ctx context.Context, userId uint) ([]domain.PostListResponse, error) {
	posts, err := r.postDAO.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var responses []domain.PostListResponse
	for _, post := range posts {
		responses = append(responses, domain.PostListResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserID:    post.UserID,
			Author:    post.User.Username,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return responses, nil
}

// Update 更新文章
func (r *PostRepository) Update(ctx context.Context, id uint, authorId uint, req domain.UpdatePostRequest) error {
	return r.postDAO.Update(ctx, id, authorId, req.Title, req.Content)
}

// Delete 删除文章
func (r *PostRepository) Delete(ctx context.Context, id uint, authorId uint) error {
	return r.postDAO.Delete(ctx, id, authorId)
}

// IsAuthor 检查用户是否为文章作者
func (r *PostRepository) IsAuthor(ctx context.Context, postId uint, userId uint) (bool, error) {
	return r.postDAO.IsAuthor(ctx, postId, userId)
}

// Exists 检查文章是否存在
func (r *PostRepository) Exists(ctx context.Context, id uint) (bool, error) {
	return r.postDAO.Exists(ctx, id)
}
