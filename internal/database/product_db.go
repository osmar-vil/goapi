package database

import (
	"database/sql"

	"github.com/osmar-vil/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db}
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	categoryErr := pd.db.QueryRow("SELECT * FROM categories WHERE id = ?", &product.CategoryID).Err()
	if categoryErr != nil {
		return nil, categoryErr
	}

	_, err := pd.db.Exec("INSERT INTO products (id, name, description, category_id, price, image_url), (?,?,?,?,?,?)",
		&product.ID, &product.Name, &product.Description, &product.CategoryID, &product.Price, &product.ImageURL)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, category_id, price, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []*entity.Product
	for rows.Next() {
		var item entity.Product
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CategoryID, &item.Price, &item.ImageURL); err != nil {
			return nil, err
		}
		response = append(response, &item)
	}

	return response, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var response entity.Product
	err := pd.db.QueryRow("SELECT id, name, description, category_id, price, image_url FROM products WHERE id = ?", id).
		Scan(&response.ID, &response.Name, &response.Description, &response.CategoryID, &response.Price, &response.ImageURL)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (pd *ProductDB) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, e := pd.db.Query("SELECT id, name, description, category_id, price, image_url FROM products WHERE category_id = ?", categoryID)

	if e != nil {
		return nil, e
	}
	defer rows.Close()

	var response []*entity.Product
	for rows.Next() {
		var item entity.Product
		if e := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CategoryID, &item.Price, &item.ImageURL); e != nil {
			return nil, e
		}
		response = append(response, &item)
	}
	return response, nil
}
