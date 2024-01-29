package database

import (
	"database/sql"

	"github.com/osmar-vil/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []*entity.Category
	for rows.Next() {
		var item entity.Category
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		response = append(response, &item)
	}

	return response, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var response entity.Category
	err := cd.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&response.ID, &response.Name)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (*entity.Category, error) {
	_, err := cd.db.Exec("INSERT INTO categories (id, name) VALUES (?,?)", category.ID, category.Name)

	if err != nil {
		return nil, err
	}

	return category, nil
}
