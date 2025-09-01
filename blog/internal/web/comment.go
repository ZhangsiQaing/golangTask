package web

import (
	"blog/internal/domain"
	"blog/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) RegisterRoutes(server *gin.Engine) {
	comments := server.Group("/comments")
	{
		comments.GET("/post/:postId", h.GetCommentsByPostId) // 获取某篇文章的所有评论
		comments.POST("/create", h.CreateComment)            // 创建评论（需要认证）
	}
}

// GetCommentsByPostId 获取某篇文章的所有评论
func (h *CommentHandler) GetCommentsByPostId(ctx *gin.Context) {
	postIdStr := ctx.Param("postId")
	postId, err := strconv.ParseUint(postIdStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusOK, ParamError)
		return
	}

	comments, err := h.commentService.GetCommentsByPostId(ctx, uint(postId))
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "",
			Message:   "获取评论列表成功",
			Data:      comments,
		})
	case service.ErrInvalidPostID, service.ErrPostNotFound:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "post_not_found",
			Message:   "文章不存在",
		})
	default:
		ctx.JSON(http.StatusOK, SystemError)
	}
}

// CreateComment 创建评论
func (h *CommentHandler) CreateComment(ctx *gin.Context) {
	var req domain.CreateCommentRequest
	if err := ctx.Bind(&req); err != nil {
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
	err := h.commentService.CreateComment(ctx, req, authorID)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "",
			Message:   "评论创建成功",
		})
	case service.ErrEmptyContent:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "param_error",
			Message:   "评论内容不能为空",
		})
	case service.ErrInvalidPostID, service.ErrPostNotFound:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "post_not_found",
			Message:   "文章不存在",
		})
	default:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "comment_create_failed",
			Message:   "创建评论失败",
		})
	}
}
