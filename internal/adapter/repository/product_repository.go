package repository

import (
	entity "go-api/internal/core/domain"
)

// ProductRepository signature of methods
type ProductRepository interface {
	GetProducts() ([]entity.Product, error)
	CreateProduct(product entity.Product) (int, error)
	GetProductById(idProduct int) (*entity.Product, error)
	DeleteProduct(idProduct int) error
	UpdateProduct(idProduct int, product entity.Product) error
}
