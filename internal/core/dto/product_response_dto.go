package dto

import entity "go-api/internal/core/domain"

type ProductResponseBody struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *ProductResponseBody) ParseFromEntities(product entity.Product) *ProductResponseBody {
	return &ProductResponseBody{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}
