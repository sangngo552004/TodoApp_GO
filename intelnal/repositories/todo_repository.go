package repositories

import (
	"awesomeProject1/intelnal/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	CreateTodo(todo *models.Todo) error
	UpdateTodo(todo *models.Todo) error
	FindByID(id uint) (*models.Todo, error)
	DeleteTodo(id uint) error
	FindByUserId(id uint) ([]models.Todo, error)
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

func (r *TodoRepositoryImpl) UpdateTodo(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepositoryImpl) FindByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *TodoRepositoryImpl) DeleteTodo(id uint) error {
	return r.db.Delete(&models.Todo{}, id).Error
}

func (r *TodoRepositoryImpl) FindByUserId(userId uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where("user_id = ?", userId).Find(&todos).Error
	return todos, err
}
