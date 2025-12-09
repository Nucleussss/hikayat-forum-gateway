package transport

import (
	"github.com/Nucleussss/hikayat-forum-gateway/internal/handler"
	"github.com/Nucleussss/hikayat-forum-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler, postHandler *handler.PostHandler) *gin.Engine {
	r := gin.Default()

	// public routes
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	// protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// auth routes
		protected.PATCH("/auth/profile", authHandler.UpdateUserProfile)
		protected.PATCH("/auth/profile/email", authHandler.ChangeUserEmail)
		protected.PATCH("/auth/profile/password", authHandler.ChangeUserPassword)
		protected.DELETE("/auth/profile", authHandler.DeleteUser)

		// post routes
		protected.POST("/posts", postHandler.CreatePost)
		protected.GET("/posts:id", postHandler.GetPost)
		protected.GET("/posts", postHandler.ListPost)
		protected.PATCH("/posts:id", postHandler.UpdatePost)
		protected.DELETE("/post:id", postHandler.DeletePost)
	}

	return r
}
