package routes

import (
	"awesomeProject1/intelnal/handlers"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
	v1 := r.Group("/api/v1")
	{
		todos := v1.Group("/auth")
		{
			todos.POST("login", authHandler.Login)
			todos.POST("register", authHandler.Register)
		}
	}
}
