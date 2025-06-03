package dto

import entity "go-api/internal/core/domain"

type ProductCreateRequestBody struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductRequestParam struct {
	Id int `json:"product_id"`
}

func (p *ProductCreateRequestBody) ParseFromEntities(product entity.Product) *ProductCreateRequestBody {
	return &ProductCreateRequestBody{
		Name:  product.Name,
		Price: product.Price,
	}
}

func (p *ProductRequestParam) ParseFromEntities(product entity.Product) *ProductRequestParam {
	return &ProductRequestParam{
		Id: product.ID,
	}
}
