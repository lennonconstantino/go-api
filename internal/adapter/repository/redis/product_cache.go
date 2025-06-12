package redis

import (
	"encoding/json"
	"fmt"
	"go-api/internal/adapter/repository"
	entity "go-api/internal/core/domain"
)

// ProductCacheRepository signature of methods
type ProductCacheRepository interface {
	GetProducts() ([]entity.Product, error)
	GetProductById(idProduct int) (*entity.Product, error)
}

// ProductCacheRepositoryImpl implements the methods
type ProductCacheRepositoryImpl struct {
	cache   CacheRepository
	product repository.ProductRepository
}

// NewProductCacheRepository initialize the repo
func NewProductCacheRepository(cache CacheRepository, product repository.ProductRepository) *ProductCacheRepositoryImpl {
	return &ProductCacheRepositoryImpl{
		cache:   cache,
		product: product,
	}
}

func (pr ProductCacheRepositoryImpl) GetProducts() ([]entity.Product, error) {
	var products []entity.Product

	reply, err := pr.cache.Get("products")
	if err != nil {
		products, err = pr.product.GetProducts()
		if err != nil {
			return nil, err
		}
		productBytes, _ := json.Marshal(products)
		pr.cache.Set("products", productBytes, 10)

		return products, nil
	}

	json.Unmarshal(reply, &products)
	return products, nil
}

func (pr ProductCacheRepositoryImpl) GetProductById(idProduct int) (*entity.Product, error) {
	var product *entity.Product

	reply, err := pr.cache.Get(fmt.Sprintf("products:%d", idProduct))
	if err != nil {
		product, err = pr.product.GetProductById(idProduct)
		if err != nil {
			return nil, err
		}
		productBytes, _ := json.Marshal(product)
		pr.cache.Set(fmt.Sprintf("products:%d", idProduct), productBytes, 10)

		return product, nil
	}

	json.Unmarshal(reply, &product)
	return product, nil
}
