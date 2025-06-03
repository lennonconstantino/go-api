package usecase

import (
	"go-api/internal/adapter/repository"
	entity "go-api/internal/core/domain"
)

// ProductUsecase signature of methods
type ProductUsecase interface {
	GetProducts() ([]entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
	GetProductById(idProduct int) (*entity.Product, error)
	DeleteProduct(idProduct int) error
	UpdateProduct(idProduct int, product entity.Product) error
}

// ProductUsecaseImpl implements the methods
type ProductUsecaseImpl struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecaseImpl {
	return &ProductUsecaseImpl{
		repository: repo,
	}
}

func (pu ProductUsecaseImpl) GetProducts() ([]entity.Product, error) {
	return pu.repository.GetProducts()
}

func (pu ProductUsecaseImpl) CreateProduct(product entity.Product) (entity.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return entity.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu ProductUsecaseImpl) GetProductById(idProduct int) (*entity.Product, error) {
	return pu.repository.GetProductById(idProduct)
}

func (pu ProductUsecaseImpl) DeleteProduct(idProduct int) error {
	return pu.repository.DeleteProduct(idProduct)
}

func (pu ProductUsecaseImpl) UpdateProduct(idProduct int, product entity.Product) error {
	return pu.repository.UpdateProduct(idProduct, product)
}
