package service

import (
	"github.com/osmar-vil/goapi/internal/database"
	"github.com/osmar-vil/goapi/internal/entity"
)

type CategoryService struct {
	categoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{categoryDB}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, e := cs.categoryDB.GetCategories()
	if e != nil {
		return nil, e
	}
	return categories, nil
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	newEntity := entity.NewCategory(name)

	response, e := cs.categoryDB.CreateCategory(newEntity)
	if e != nil {
		return nil, e
	}
	return response, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	response, e := cs.categoryDB.GetCategory(id)
	if e != nil {
		return nil, e
	}
	return response, nil
}
