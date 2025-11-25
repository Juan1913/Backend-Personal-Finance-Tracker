package errors

import (
	authMiddleware "apiGo/internal/auth/middleware"
	"apiGo/internal/user/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	// Documentación Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		usersGroup := api.Group("/user")
		{
			usersGroup.POST("", userHandler.Create)
			usersGroup.GET("", authMiddleware.RequireRoles("admin"), userHandler.GetAll)
			usersGroup.GET(":id", authMiddleware.RequireRoles("admin", "user"), userHandler.GetByID)
			usersGroup.PUT(":id", authMiddleware.RequireRoles("admin"), userHandler.Update)
			usersGroup.DELETE(":id", authMiddleware.RequireRoles("admin"), userHandler.Delete)
		}
	}
}
