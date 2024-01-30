package service

import (
	"github.com/osmar-vil/goapi/internal/database"
	"github.com/osmar-vil/goapi/internal/entity"
)

type ProductService struct {
	productDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{productDB}
}

func (ps *ProductService) CreateProduct(
	name, description, image_url, category_id string,
	price float64,
) (*entity.Product, error) {
	newEntity := entity.NewProduct(name, description, category_id, image_url, price)

	_, e := ps.productDB.CreateProduct(newEntity)
	if e != nil {
		return nil, e
	}
	return newEntity, e
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	response, e := ps.productDB.GetProducts()
	if e != nil {
		return nil, e
	}
	return response, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	response, e := ps.productDB.GetProduct(id)
	if e != nil {
		return nil, e
	}
	return response, nil
}

func (ps *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	response, e := ps.productDB.GetProductsByCategoryID(categoryID)
	if e != nil {
		return nil, e
	}
	return response, nil
}
