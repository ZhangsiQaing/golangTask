package web

type Response struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
}

var (
	// 系统错误
	SystemError = Response{
		ErrorCode: "system_error",
		Message:   "系统错误，请稍后再试",
	}

	// 参数错误
	ParamError = Response{
		ErrorCode: "param_error",
		Message:   "参数错误",
	}

	// 文章相关错误
	PostNotFoundError = Response{
		ErrorCode: "post_not_found",
		Message:   "文章不存在",
	}

	PostPermissionError = Response{
		ErrorCode: "post_permission_denied",
		Message:   "权限不足，只有文章作者才能操作",
	}

	PostCreateError = Response{
		ErrorCode: "post_create_failed",
		Message:   "创建文章失败",
	}

	PostUpdateError = Response{
		ErrorCode: "post_update_failed",
		Message:   "更新文章失败",
	}

	PostDeleteError = Response{
		ErrorCode: "post_delete_failed",
		Message:   "删除文章失败",
	}

	PostListError = Response{
		ErrorCode: "post_list_failed",
		Message:   "获取文章列表失败",
	}

	// 评论相关错误
	CommentNotFoundError = Response{
		ErrorCode: "comment_not_found",
		Message:   "评论不存在",
	}

	CommentPermissionError = Response{
		ErrorCode: "comment_permission_denied",
		Message:   "权限不足，只有评论作者才能操作",
	}

	CommentCreateError = Response{
		ErrorCode: "comment_create_failed",
		Message:   "创建评论失败",
	}

	CommentUpdateError = Response{
		ErrorCode: "comment_update_failed",
		Message:   "更新评论失败",
	}

	CommentDeleteError = Response{
		ErrorCode: "comment_delete_failed",
		Message:   "删除评论失败",
	}

	CommentListError = Response{
		ErrorCode: "comment_list_failed",
		Message:   "获取评论列表失败",
	}
)
