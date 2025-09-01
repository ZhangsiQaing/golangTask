//go:build wireinject

package main

import (
	"blog/internal/dao"
	"blog/internal/ioc"
	"blog/internal/repository"
	"blog/internal/service"
	"blog/internal/web"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitDB,
		ioc.InitLogger,

		dao.ProviderSet,
		repository.ProviderSet,
		service.ProviderSet,
		web.ProviderSet,

		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
