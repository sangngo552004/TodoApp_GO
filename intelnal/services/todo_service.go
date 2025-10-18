package services

import (
	"awesomeProject1/intelnal/apperror"
	"awesomeProject1/intelnal/dtos/dto_requests"
	"awesomeProject1/intelnal/models"
	"awesomeProject1/intelnal/repositories"
)

type TodoService interface {
	GetTodos(userId uint) ([]models.Todo, *apperror.AppError)
	CreateTodo(userId uint, req *dto_requests.TodoRequest) (*models.Todo, *apperror.AppError)
	UpdateTodo(userId uint, id uint, req *dto_requests.TodoRequest) (*models.Todo, *apperror.AppError)
	DeleteTodo(userId uint, id uint) *apperror.AppError
}

type TodoServiceImpl struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepository: todoRepository}
}

func (s *TodoServiceImpl) GetTodos(userId uint) ([]models.Todo, *apperror.AppError) {
	todos, err := s.todoRepository.FindByUserId(userId)
	if err != nil {
		return nil, apperror.InternalServerError("Failed to get todos", err)
	}
	return todos, nil
}

func (s *TodoServiceImpl) CreateTodo(userId uint, req *dto_requests.TodoRequest) (*models.Todo, *apperror.AppError) {
	todo := &models.Todo{
		Title:  req.Title,
		UserId: userId,
	}
	if err := s.todoRepository.CreateTodo(todo); err != nil {
		return nil, apperror.InternalServerError("Failed to create todo", err)
	}
	return todo, nil
}

func (s *TodoServiceImpl) UpdateTodo(userId uint, id uint, req *dto_requests.TodoRequest) (*models.Todo, *apperror.AppError) {
	todo, err := s.todoRepository.FindByID(id)
	if err != nil {
		return nil, apperror.NotFound("Todo not found", err)
	}

	if todo.UserId != userId {
		return nil, apperror.Forbidden("You are not authorized to update this todo", nil)
	}

	todo.Title = req.Title
	todo.Completed = req.Completed
	if err := s.todoRepository.UpdateTodo(todo); err != nil {
		return nil, apperror.InternalServerError("Failed to update todo", err)
	}
	return todo, nil
}

func (s *TodoServiceImpl) DeleteTodo(userId uint, id uint) *apperror.AppError {
	todo, err := s.todoRepository.FindByID(id)
	if err != nil {
		return apperror.NotFound("Todo not found", err)
	}

	if todo.UserId != userId {
		return apperror.Forbidden("You are not authorized to delete this todo", nil)
	}

	if err := s.todoRepository.DeleteTodo(id); err != nil {
		return apperror.InternalServerError("Failed to delete todo", err)
	}
	return nil
}
