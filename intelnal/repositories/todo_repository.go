package repositories

import (
	"awesomeProject1/intelnal/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	CreateTodo(todo *models.Todo) error
}

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepositoryImpl) CreateTodo(todo *models.Todo) error {
	return r.db.Create(todo).Error
}
