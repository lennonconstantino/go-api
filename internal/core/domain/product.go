package entity

type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (t *Product) TableName() string {
	return "products"
}
