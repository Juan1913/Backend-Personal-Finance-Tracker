package account

import (
	"apiGo/internal/account/handler"
	authMiddleware "apiGo/internal/auth/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(r *gin.Engine, accountHandler *handler.AccountHandler) {
	api := r.Group("/api")
	{
		accountsGroup := api.Group("/accounts")
		{
			accountsGroup.POST("", authMiddleware.RequireRoles("admin", "user"), accountHandler.Create)
			accountsGroup.GET("", authMiddleware.RequireRoles("admin"), accountHandler.GetAll)
			accountsGroup.GET("/user/:userId", authMiddleware.RequireRoles("admin", "user"), accountHandler.GetByUserID)
			accountsGroup.GET("/:id", authMiddleware.RequireRoles("admin", "user"), accountHandler.GetByID)
			accountsGroup.PUT("/:id", authMiddleware.RequireRoles("admin", "user"), accountHandler.Update)
			accountsGroup.DELETE("/:id", authMiddleware.RequireRoles("admin", "user"), accountHandler.Delete)
		}
	}
}
