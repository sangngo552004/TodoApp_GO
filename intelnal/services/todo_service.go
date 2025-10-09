package services

import (
	"awesomeProject1/intelnal/models"
	"awesomeProject1/intelnal/repositories"
	"errors"
)

type TodoService interface {
	GetTodos(userId uint) ([]models.Todo, error)
	CreateTodo(userId uint, todo *models.Todo) error
	UpdateTodo(userId uint, id uint, completed bool) error
	DeleteTodo(userId uint, id uint) error
}

type TodoServiceImpl struct {
	todoRepository repositories.TodoRepository
	userRepository repositories.UserRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepository: todoRepository}
}

func (s *TodoServiceImpl) GetTodos(userId uint) ([]models.Todo, error) {
	return s.todoRepository.FindByUserId(userId)
}

func (s *TodoServiceImpl) CreateTodo(userId uint, todo *models.Todo) error {
	todo.UserId = userId
	return s.todoRepository.CreateTodo(todo)
}

func (s *TodoServiceImpl) UpdateTodo(userId uint, id uint, completed bool) error {
	todo, err := s.todoRepository.FindByID(id)
	if err != nil {
		return err
	}
	if todo.UserId != userId {
		return errors.New("unauthorized to update this todo")
	}
	todo.Completed = completed
	return s.todoRepository.UpdateTodo(todo)
}

func (s *TodoServiceImpl) DeleteTodo(userId uint, id uint) error {
	todo, err := s.todoRepository.FindByID(id)
	if err != nil {
		return err
	}
	if todo.UserId != userId {
		return errors.New("unauthorized to delete this todo")
	}
	return s.todoRepository.DeleteTodo(id)

}
