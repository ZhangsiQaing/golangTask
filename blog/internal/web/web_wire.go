package web

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserHandler,
	NewPostHandler,
	NewCommentHandler,
)
