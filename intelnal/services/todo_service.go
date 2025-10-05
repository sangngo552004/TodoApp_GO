package services

import (
	"awesomeProject1/intelnal/models"
	"awesomeProject1/intelnal/repositories"
)

type TodoService interface {
	GetTodos() ([]models.Todo, error)
}

type TodoServiceImpl struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepository: todoRepository}
}

func (s *TodoServiceImpl) GetTodos() ([]models.Todo, error) {
	return s.todoRepository.GetAll()
}
