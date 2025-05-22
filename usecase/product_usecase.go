package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase interface {
	GetProducts() ([]model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	GetProductById(id_product int) (*model.Product, error)
	DeleteProduct(id_product int) error
	UpdateProduct(id_product int, product model.Product) error
}

type ProductUsecaseImpl struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecaseImpl {
	return &ProductUsecaseImpl{
		repository: repo,
	}
}

func (pu ProductUsecaseImpl) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu ProductUsecaseImpl) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu ProductUsecaseImpl) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu ProductUsecaseImpl) DeleteProduct(id_product int) error {
	if err := pu.repository.DeleteProduct(id_product); err != nil {
		return err
	}
	return nil
}

func (pu ProductUsecaseImpl) UpdateProduct(id_product int, product model.Product) error {
	if err := pu.repository.UpdateProduct(id_product, product); err != nil {
		return err
	}
	return nil
}
