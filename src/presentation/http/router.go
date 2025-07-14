package http

import (
	"github.com/eyagovbusiness/GSWB.Users/src/presentation/http/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.ListUsers)
		}
	}

	return router
}
