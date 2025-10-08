package routes

import (
	"awesomeProject1/intelnal/handlers"
	"awesomeProject1/intelnal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(r *gin.Engine, todoHandler *handlers.TodoHandler) {
	v1 := r.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		todos.Use(middlewares.JWTAuthMiddleware())
		{
			todos.GET("", todoHandler.GetTodos)
			todos.POST("", todoHandler.CreateTodo)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}
	}
}
