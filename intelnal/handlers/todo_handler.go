package handlers

import (
	"awesomeProject1/intelnal/dtos/dto_requests"
	"awesomeProject1/intelnal/dtos/dto_responses"
	"awesomeProject1/intelnal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service services.TodoService
}

func NewTodoHandler(service services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	userId := c.GetInt("user_id")
	todos, err := h.service.GetTodos(uint(userId))
	if err != nil {
		c.Error(err)
		return
	}
	todoResponse := make([]dto_responses.TodoResponse, len(todos))
	for i, todo := range todos {
		todoResponse[i].ID = todo.ID
		todoResponse[i].Title = todo.Title
		todoResponse[i].Completed = todo.Completed
	}
	dto_responses.SuccessResponse(c, http.StatusOK, "Todos fetched", todoResponse)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	userId := c.GetInt("user_id")
	var req dto_requests.TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiError := &dto_responses.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		dto_responses.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", apiError)
		return
	}

	todo, err := h.service.CreateTodo(uint(userId), &req)
	if err != nil {
		c.Error(err)
		return
	}
	todoResponse := dto_responses.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}
	dto_responses.SuccessResponse(c, http.StatusCreated, "Create todo success", todoResponse)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	userId := c.GetInt("user_id")
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		apiError := &dto_responses.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		dto_responses.ErrorResponse(c, http.StatusBadRequest, "Invalid Type of ID param", apiError)
		return
	}
	var req dto_requests.TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiError := &dto_responses.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		dto_responses.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", apiError)
		return
	}
	todo, err := h.service.UpdateTodo(uint(userId), uint(id), &req)
	if err != nil {
		c.Error(err)
	}
	todoResponse := dto_responses.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}
	dto_responses.SuccessResponse(c, http.StatusOK, "Update todo success", todoResponse)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	userId := c.GetInt("user_id")
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		apiError := &dto_responses.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		dto_responses.ErrorResponse(c, http.StatusBadRequest, "Invalid Type of ID param", apiError)
		return
	}
	if err := h.service.DeleteTodo(uint(userId), uint(id)); err != nil {
		c.Error(err)
		return
	}
	dto_responses.SuccessResponse(c, http.StatusOK, "Delete todo success", nil)
}
