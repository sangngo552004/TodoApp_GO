package routes

import (
	"awesomeProject1/intelnal/handlers"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
	v1 := r.Group("/api/v1")
	{
		auths := v1.Group("/auth")
		{
			auths.POST("login", authHandler.Login)
			auths.POST("register", authHandler.Register)
			auths.POST("refresh", authHandler.Refresh)
		}
	}
}
