package handlers

import (
	"awesomeProject1/intelnal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service services.TodoService
}

func NewTodoHandler(service services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.service.GetTodos()
	if err != nil {
		log.Println("loi o day")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}
