package postgres

import (
	"database/sql"
	"fmt"
	entity "go-api/internal/core/domain"
)

// ProductRepositoryImpl implements the methods
type ProductRepositoryImpl struct {
	connection *sql.DB
}

// NewProductRepository initialize the repo
func NewProductRepository(connection *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		connection: connection,
	}
}

// GetProducts
func (pr ProductRepositoryImpl) GetProducts() ([]entity.Product, error) {
	query := "SELECT id, name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []entity.Product{}, err
	}
	defer rows.Close()

	var productList []entity.Product
	var productObj entity.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []entity.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

// CreateProduct
func (pr ProductRepositoryImpl) CreateProduct(product entity.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(name, price)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

// GetProductById
func (pr ProductRepositoryImpl) GetProductById(idProduct int) (*entity.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	var product entity.Product

	err = query.QueryRow(idProduct).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}

// DeleteProduct
func (pr ProductRepositoryImpl) DeleteProduct(idProduct int) error {
	statement, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(idProduct); err != nil {
		return err
	}

	return nil
}

// UpdateProduct
func (pr ProductRepositoryImpl) UpdateProduct(idProduct int, product entity.Product) error {
	statement, err := pr.connection.Prepare(
		"update product set price = $1, name = $2 where id = $3",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(product.Price, product.Name, idProduct); err != nil {
		return err
	}

	return nil
}
