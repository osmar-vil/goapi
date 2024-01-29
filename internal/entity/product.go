package entity

import "github.com/google/uuid"

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryID  string
	ImageURL    string
}

func NewProduct(
	name,
	description,
	categoryId,
	imageUrl string,
	price float64,
) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryId,
		ImageURL:    imageUrl,
	}
}
