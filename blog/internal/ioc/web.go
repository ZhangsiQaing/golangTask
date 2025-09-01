package ioc

import (
	"blog/internal/middleware"
	"blog/internal/web"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitWebServer(mdls []gin.HandlerFunc, userHdl *web.UserHandler,
	postHdl *web.PostHandler,
	commentHdl *web.CommentHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	userHdl.RegisterRoutes(server)
	postHdl.RegisterRoutes(server)
	commentHdl.RegisterRoutes(server)
	return server
}

func InitGinMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		cors.New(cors.Config{
			//AllowAllOrigins: true,
			//AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,
			AllowHeaders:     []string{"Content-Type", "Authorization"},
			// 这个是允许前端访问你的后端响应中带的头部
			ExposeHeaders: []string{"x-jwt-token"},
			AllowOriginFunc: func(origin string) bool {
				// if strings.HasPrefix(origin, "http://localhost") {
				// 	//if strings.Contains(origin, "localhost") {
				// 	return true
				// }
				return strings.Contains(origin, "your_company.com")
			},
			MaxAge: 12 * time.Hour,
		}),
		// func(ctx *gin.Context) {
		// 	println("这是我的 Middleware")
		// },
		// ratelimit.NewBuilder(redisClient, time.Second, 1000).Build(),
		(&middleware.LoginJWTMiddlewareBuilder{}).CheckLogin(),
	}
}
