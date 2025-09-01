package web

import (
	"blog/internal/domain"
	"blog/internal/service"
	"blog/pkg/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
	l           logger.LoggerV1
}

func NewPostHandler(postService *service.PostService, l logger.LoggerV1) *PostHandler {
	return &PostHandler{
		postService: postService,
		l:           l,
	}
}

func (h *PostHandler) RegisterRoutes(server *gin.Engine) {
	posts := server.Group("/posts")
	{
		posts.GET("/list", h.GetAllPosts)              // 获取所有文章列表
		posts.GET("/:id", h.GetPost)                   // 获取单个文章详情
		posts.GET("/user/:userId", h.GetPostsByUserId) // 获取指定用户的文章
		posts.POST("/create", h.CreatePost)            // 创建文章（需要认证）
		posts.PUT("/:id", h.UpdatePost)                // 更新文章（需要认证且是作者）
		posts.DELETE("/:id", h.DeletePost)             // 删除文章（需要认证且是作者）
	}
}

// GetAllPosts 获取所有文章列表
func (h *PostHandler) GetAllPosts(ctx *gin.Context) {
	posts, err := h.postService.GetAllPosts(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, PostListError)
		return
	}
	h.l.Info("获取文章列表成功", logger.Field{Key: "posts", Val: posts})
	ctx.JSON(http.StatusOK, Response{
		ErrorCode: "",
		Message:   "获取文章列表成功",
		Data:      posts,
	})
}

// GetPost 获取单个文章详情
func (h *PostHandler) GetPost(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, ParamError)
		return
	}

	fmt.Println(id, err)
	// fmt.Println(userID, exists)
	post, err := h.postService.GetPost(ctx, uint(id))
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "",
			Message:   "获取文章成功",
			Data:      post,
		})
	case service.ErrPostNotFound:
		ctx.JSON(http.StatusOK, PostNotFoundError)
	default:
		ctx.JSON(http.StatusOK, SystemError)
	}
}

// GetPostsByUserId 获取指定用户的文章列表
func (h *PostHandler) GetPostsByUserId(ctx *gin.Context) {
	userIdStr := ctx.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, ParamError)
		return
	}

	posts, err := h.postService.GetPostsByUserId(ctx, uint(userId))
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "",
			Message:   "获取用户文章成功",
			Data:      posts,
		})
	case service.ErrInvalidUserID:
		ctx.JSON(http.StatusOK, ParamError)
	default:
		ctx.JSON(http.StatusOK, SystemError)
	}
}

// CreatePost 创建文章
func (h *PostHandler) CreatePost(ctx *gin.Context) {
	var req domain.CreatePostRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusOK, ParamError)
		return
	}

	// 从JWT中获取用户ID（这里假设已经通过中间件设置）
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "unauthorized",
			Message:   "用户未认证",
		})
		return
	}

	authorID := uint(userID.(int64))
	err := h.postService.CreatePost(ctx, req, authorID)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "",
			Message:   "文章创建成功",
		})
	case service.ErrEmptyTitle, service.ErrEmptyContent, service.ErrTitleTooLong:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "param_error",
			Message:   err.Error(),
		})
	default:
		ctx.JSON(http.StatusOK, PostCreateError)
	}
}

// UpdatePost 更新文章
func (h *PostHandler) UpdatePost(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, ParamError)
		return
	}

	var req domain.UpdatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, ParamError)
		return
	}

	// 从JWT中获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "unauthorized",
			Message:   "用户未认证",
		})
		return
	}

	authorID := uint(userID.(int64))
	err = h.postService.UpdatePost(ctx, uint(id), authorID, req)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "",
			Message:   "文章更新成功",
		})
	case service.ErrPostNotFound:
		ctx.JSON(http.StatusOK, PostNotFoundError)
	case service.ErrPostPermission:
		ctx.JSON(http.StatusOK, PostPermissionError)
	case service.ErrEmptyTitle, service.ErrEmptyContent, service.ErrTitleTooLong:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "param_error",
			Message:   err.Error(),
		})
	default:
		ctx.JSON(http.StatusOK, PostUpdateError)
	}
}

// DeletePost 删除文章
func (h *PostHandler) DeletePost(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, ParamError)
		return
	}

	// 从JWT中获取用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "unauthorized",
			Message:   "用户未认证",
		})
		return
	}

	authorID := uint(userID.(int64))
	err = h.postService.DeletePost(ctx, uint(id), authorID)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "",
			Message:   "文章删除成功",
		})
	case service.ErrPostNotFound:
		ctx.JSON(http.StatusOK, PostNotFoundError)
	case service.ErrPostPermission:
		ctx.JSON(http.StatusOK, PostPermissionError)
	default:
		ctx.JSON(http.StatusOK, PostDeleteError)
	}
}
