package postgres

import (
	"golang-practice/pkg/todo"

	"gorm.io/gorm"
)

type RepositoryStatePG struct {
	db *gorm.DB
}

func New(db *gorm.DB) todo.TodoRepository {
	return &RepositoryStatePG{
		db: db,
	}
}

func (t *RepositoryStatePG) Save(todoTask todo.TodoTask) error {
	return nil
}

func (t *RepositoryStatePG) GetById(id string) (todo.TodoTask, error) {
	return nil, nil
}

func (t *RepositoryStatePG) GetAll() ([]todo.TodoTask, error) {
	return nil, nil
}
