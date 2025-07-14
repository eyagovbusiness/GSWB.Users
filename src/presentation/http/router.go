package http

import (
	"github.com/eyagovbusiness/GSWB.Users/src/presentation/http/handler"
	"github.com/eyagovbusiness/GSWB.Users/src/presentation/http/middleware"

	_ "github.com/eyagovbusiness/GSWB.Users/docs" // <-- required for generated Swagger docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.New() // note: New() does NOT include default middleware

	router.Use(gin.Recovery())             // catch panics
	router.Use(middleware.RequestLogger()) // our custom logging

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.ListUsers)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
