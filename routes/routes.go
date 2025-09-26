package routes

import (
	"golang_blog/controllers"
	"golang_blog/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes() *gin.Engine {
	r := gin.New()

	// 使用中间件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandlerMiddleware())
	r.Use(gin.Recovery())

	// 创建控制器实例
	authController := &controllers.AuthController{}
	postController := &controllers.PostController{}
	commentController := &controllers.CommentController{}

	// API路由组
	api := r.Group("/api/v1")
	{
		// 认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}

		// 需要认证的路由
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware())
		{
			// 用户信息
			authenticated.GET("/profile", authController.GetProfile)

			// 文章相关路由
			posts := authenticated.Group("/posts")
			{
				posts.POST("", postController.CreatePost)
				posts.PUT("/:id", postController.UpdatePost)
				posts.DELETE("/:id", postController.DeletePost)
			}

			// 评论相关路由
			comments := authenticated.Group("/posts/:post_id/comments")
			{
				comments.POST("", commentController.CreateComment)
			}
		}

		// 公开路由（无需认证）
		public := api.Group("")
		{
			// 文章公开路由
			public.GET("/posts", postController.GetPosts)
			public.GET("/posts/:id", postController.GetPost)
		}

		// 评论公开路由（单独分组避免路由冲突）
		comments := api.Group("/comments")
		{
			comments.GET("/post/:post_id", commentController.GetComments)
		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Blog API is running",
		})
	})

	return r
}